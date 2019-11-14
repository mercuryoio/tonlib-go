package tonlib

//func TestClient_CreatePrivateKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	_, err = cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//}
//
//func TestClient_DeletePrivateKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	if err = cln.DeletePrivateKey(pkey, []byte(TEST_PASSWORD)); err != nil {
//		t.Fatal("Ton delete key error", err)
//	}
//}
//
//func TestClient_ExportPrivateKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	_, err = cln.ExportPrivateKey(pkey, []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton export private key error", err)
//	}
//}
//
//func TestClient_ImportKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	mnemonics, err := cln.ExportPrivateKey(pkey, []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton export private key error", err)
//	}
//	if _, err := cln.ImportKey(mnemonics, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD)); err != nil {
//		t.Fatal("Ton import private key error", err)
//	}
//}
//
//func TestClient_ExportPemKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	if pem, err := cln.ExportPemKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD)); err != nil || len(pem) == 0 {
//		t.Fatal("Ton export pem key error", err)
//	}
//}
//
//func TestClient_ExportEncryptedKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	if _, err = cln.ExportEncryptedKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD)); err != nil {
//		t.Fatal("Ton export pem key error", err)
//	}
//}
//
//func TestClient_ChangeLocalPassword(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key for change password error", err)
//	}
//	_, err = cln.ChangeLocalPassword(pKey, []byte(TEST_PASSWORD), []byte(TEST_CHANGE_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton change key password error", err)
//	}
//}
//
//func TestClient_ImportPemKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	key, err := cln.ExportPemKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton export pem key error", err)
//	}
//
//	_, err = cln.ImportPemKey(key, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton import pem key error", err)
//	}
//}
//
//func TestClient_ImportEncryptedKey(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key error", err)
//	}
//	key, err := cln.ExportEncryptedKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton export encrypted key error", err)
//	}
//
//	_, err = cln.ImportEncryptedKey(key, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton import encrypted key error", err)
//	}
//}
//
