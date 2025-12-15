package hades

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"simple", "simple@example.com", true},
		{"plus_tag", "user.name+tag@example.co", true},
		{"underscore_and_country", "user_name@example.co.uk", true},
		{"subdomain", "user-name@sub.domain.com", true},
		{"uppercase", "USER@EXAMPLE.COM", true},
		{"double_dot_domain", "user@site..com", true},        // accepted by the simplified regex
		{"leading_dot_local", ".user@example.com", true},     // accepted by the simplified regex
		{"leading_hyphen_domain", "user@-example.com", true}, // accepted by the simplified regex

		{"no_at", "plainaddress", false},
		{"empty_local", "@missinglocal.org", false},
		{"missing_at_sign", "missingatsign.com", false},
		{"dot_only_domain", "user@.com", false},
		{"no_dot_in_domain", "user@com", false},
		{"one_letter_tld", "user@site.c", false},
		{"space_in_local", "user name@example.com", false},
		{"empty_string", "", false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValidEmail(tc.email); got != tc.want {
				t.Fatalf("IsValidEmail(%q) = %v; want %v", tc.email, got, tc.want)
			}
		})
	}
}
