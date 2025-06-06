package wireguard

import "testing"

func TestGenerateProfile(t *testing.T) {
	var expectedResult = `[Interface]
PrivateKey = 1
Address = 2/32, 3/128
DNS = 127.0.0.1
MTU = 1420

[Peer]
PublicKey = 4
AllowedIPs = 0.0.0.0/0, ::/0
Endpoint = 5
`

	result, err := generateProfile(&ProfileData{
		PrivateKey: "1",
		Address1:   "2",
		Address2:   "3",
		PublicKey:  "4",
		Endpoint:   "5",
	})
	if err != nil {
		t.Error(err)
	}

	if expectedResult != result {
		t.Error()
	}
}
