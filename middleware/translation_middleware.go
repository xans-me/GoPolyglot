package middleware

import (
	"encoding/json"
	"github.com/xans-me/GoPolyglot"
	"net/http"
)

// AgnosticTranslationMiddleware creates a middleware that handles translation for any http.Handler,
// making it framework-agnostic and suitable for both Gin and Mux.
func AgnosticTranslationMiddleware(translator *GoPolyglot.Translator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseCapture := newBufferedResponseWriter(w)
			next.ServeHTTP(responseCapture, r)

			resp, err := getResponseMap(responseCapture)
			if err != nil {
				// In case of an error in unmarshalling, write back the original captured response.
				w.WriteHeader(responseCapture.Status())
				w.Write(responseCapture.body.Bytes())
				return
			}

			// Translate the response based on the request headers.
			err = translateBasedOnHeader(resp, translator, r)
			handleTranslationError(w, err)

			// If there's no error, proceed to write the potentially modified response.
			if err == nil {
				processResponse(w, resp, responseCapture)
			}
		})
	}
}

// getResponseMap unmarshal the captured response body into a map for further processing.
func getResponseMap(responseCapture *bufferedResponseWriter) (map[string]interface{}, error) {
	var resp map[string]interface{}
	err := json.Unmarshal(responseCapture.body.Bytes(), &resp)
	return resp, err
}

// translateBasedOnHeader determines whether dynamic or static translation is needed based on the request headers.
func translateBasedOnHeader(resp map[string]interface{}, translator *GoPolyglot.Translator, r *http.Request) error {
	return translateResponse(resp, translator, r.Header.Get(GoPolyglot.AcceptedLanguageHeader))
}

// translateResponse handles static translation by updating specific response fields based on the provided language.
func translateResponse(resp map[string]interface{}, translator *GoPolyglot.Translator, language string) error {
	rc, rcOK := resp["rc"].(string)
	trxType, trxTypeOK := resp["trxType"].(string)
	trxFeature, trxFeatureOK := resp["trxFeature"].(string)

	if !rcOK || !trxTypeOK || !trxFeatureOK {
		// Skip translation if type assertion fails for any field.
		return nil
	}

	translated, err := translator.TranslateWithParams(rc, trxType, trxFeature, language)
	if err != nil {
		return err
	}

	resp["title"] = translated.Title
	resp["description"] = translated.Description
	return nil
}

// processResponse serializes the response map and writes it to the original http.ResponseWriter.
func processResponse(_ http.ResponseWriter, resp map[string]interface{}, responseCapture *bufferedResponseWriter) {
	respBytes, _ := json.Marshal(resp)
	responseCapture.ResponseWriter.WriteHeader(responseCapture.Status())
	responseCapture.ResponseWriter.Write(respBytes)
}

// handleTranslationError sends an internal server error response if translation fails.
func handleTranslationError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
