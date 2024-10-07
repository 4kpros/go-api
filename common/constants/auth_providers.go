package constants

import "slices"

const AUTH_PROVIDER_GOOGLE = "google"
const AUTH_PROVIDER_FACEBOOK = "facebook"

var AllAuthProviders = []string{
	AUTH_PROVIDER_GOOGLE,
	AUTH_PROVIDER_FACEBOOK,
}

func IsAuthProviderValid(provider string) bool {
	return slices.Contains(AllAuthProviders, provider)
}
