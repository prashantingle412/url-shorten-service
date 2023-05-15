# Techdocs for url shortner Service

## REST api

    * Enocde : This api is available on /shortner/encode endpoint. Encode api is developed to encode the long or orignial url to the shorten form and save into non persistence stoarage. Here used Map for in-memory data stoarage. It accept http POST method with JSON request body and will return JSON response body with encoded URL, message and http status code.
