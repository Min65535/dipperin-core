package test_util

type CodeAndAbi struct {
	code []byte
	abi  string
	hpp  string
	cpp  string
}


var contractTestData = map[string]CodeAndAbi{
	"event":{
		code: []byte{0,97,115,109,1,0,0,0,1,47,9,96,0,0,96,1,127,0,96,3,127,127,127,1,127,96,1,127,1,127,96,4,127,127,127,127,0,96,2,127,127,0,96,1,126,0,96,3,127,127,127,0,96,1,127,1,126,2,199,1,12,3,101,110,118,9,99,97,108,108,86,97,108,117,101,0,1,3,101,110,118,7,109,101,109,109,111,118,101,0,2,3,101,110,118,6,109,101,109,99,112,121,0,2,3,101,110,118,41,95,90,78,53,98,111,111,115,116,49,53,116,104,114,111,119,95,101,120,99,101,112,116,105,111,110,69,82,75,83,116,57,101,120,99,101,112,116,105,111,110,0,1,3,101,110,118,6,109,101,109,115,101,116,0,2,3,101,110,118,5,97,98,111,114,116,0,0,3,101,110,118,6,109,97,108,108,111,99,0,3,3,101,110,118,4,102,114,101,101,0,1,3,101,110,118,9,101,109,105,116,69,118,101,110,116,0,4,3,101,110,118,6,112,114,105,110,116,115,0,1,3,101,110,118,8,112,114,105,110,116,115,95,108,0,5,3,101,110,118,7,112,114,105,110,116,117,105,0,6,3,25,24,0,3,3,1,1,3,1,3,1,0,0,1,5,0,0,0,0,7,7,4,0,3,8,8,4,5,1,112,1,5,5,5,3,1,0,2,6,21,3,127,1,65,208,138,4,11,127,0,65,208,138,4,11,127,0,65,198,10,11,7,118,11,6,109,101,109,111,114,121,2,0,11,95,95,104,101,97,112,95,98,97,115,101,3,1,10,95,95,100,97,116,97,95,101,110,100,3,2,5,95,90,110,119,109,0,13,5,95,90,110,97,109,0,14,6,95,90,100,108,80,118,0,15,6,95,90,100,97,80,118,0,16,4,105,110,105,116,0,32,12,114,101,116,117,114,110,83,116,114,105,110,103,0,33,9,114,101,116,117,114,110,73,110,116,0,34,10,114,101,116,117,114,110,85,105,110,116,0,35,9,10,1,0,65,1,11,4,17,18,19,20,10,132,58,24,2,0,11,56,1,2,127,2,64,32,0,65,1,32,0,27,34,1,16,6,34,0,13,0,3,64,65,0,33,0,65,0,40,2,128,8,34,2,69,13,1,32,2,17,0,0,32,1,16,6,34,0,69,13,0,11,11,32,0,11,6,0,32,0,16,13,11,14,0,2,64,32,0,69,13,0,32,0,16,7,11,11,6,0,32,0,16,15,11,59,1,2,127,32,0,65,140,8,54,2,0,32,0,40,2,4,34,1,65,124,106,34,2,32,2,40,2,0,65,127,106,34,2,54,2,0,2,64,32,2,65,127,76,13,0,32,0,15,11,32,1,65,116,106,16,15,32,0,11,63,1,2,127,32,0,65,140,8,54,2,0,32,0,40,2,4,34,1,65,124,106,34,2,32,2,40,2,0,65,127,106,34,2,54,2,0,2,64,32,2,65,127,76,13,0,32,0,16,15,15,11,32,1,65,116,106,16,15,32,0,16,15,11,7,0,32,0,40,2,4,11,63,1,2,127,32,0,65,140,8,54,2,0,32,0,40,2,4,34,1,65,124,106,34,2,32,2,40,2,0,65,127,106,34,2,54,2,0,2,64,32,2,65,127,76,13,0,32,0,16,15,15,11,32,1,65,116,106,16,15,32,0,16,15,11,5,0,16,5,0,11,5,0,16,5,0,11,253,16,1,9,127,35,0,65,192,0,107,34,1,36,0,32,1,65,16,106,65,0,54,2,0,32,1,66,0,55,3,8,65,174,8,65,3,113,33,2,65,173,8,65,3,113,33,3,65,175,8,65,3,113,33,4,65,176,8,65,3,113,33,5,65,177,8,33,6,2,64,2,64,65,172,8,65,3,113,34,7,69,13,0,32,3,65,0,71,32,2,65,0,71,113,32,4,65,0,71,113,32,5,65,0,71,113,65,177,8,65,3,113,65,0,71,113,13,1,11,65,174,8,65,173,8,32,3,27,65,175,8,32,3,69,32,2,69,114,34,6,27,65,176,8,32,6,32,4,69,114,34,6,27,65,177,8,32,6,32,5,69,114,27,65,172,8,32,7,27,65,124,106,33,6,3,64,32,6,65,4,106,34,6,40,2,0,34,3,65,127,115,32,3,65,255,253,251,119,106,113,65,128,129,130,132,120,113,69,13,0,11,32,3,65,255,1,113,69,13,0,32,6,33,3,3,64,32,3,45,0,1,33,2,32,3,65,1,106,34,6,33,3,32,2,13,0,11,11,2,64,2,64,2,64,2,64,2,64,2,64,2,64,2,64,32,6,65,172,8,107,34,6,65,112,79,13,0,2,64,2,64,2,64,32,6,65,11,79,13,0,32,1,32,6,65,1,116,58,0,8,32,1,65,8,106,65,1,114,33,3,32,6,13,1,12,2,11,32,6,65,16,106,65,112,113,34,2,16,13,33,3,32,1,32,2,65,1,114,54,2,8,32,1,32,3,54,2,16,32,1,32,6,54,2,12,11,32,3,65,172,8,32,6,16,2,26,11,32,3,32,6,106,65,0,58,0,0,32,1,65,40,106,34,3,66,0,55,3,0,32,1,65,24,106,65,8,106,66,0,55,3,0,32,1,66,0,55,3,24,65,8,16,13,34,6,66,1,55,2,0,32,3,32,6,65,8,106,34,2,54,2,0,32,1,65,44,106,32,2,54,2,0,32,1,32,6,54,2,36,32,1,65,48,106,65,8,106,65,0,54,2,0,32,1,66,0,55,3,48,32,0,33,3,2,64,2,64,2,64,2,64,32,0,65,3,113,69,13,0,32,0,45,0,0,69,13,1,32,0,65,1,106,33,6,3,64,32,6,34,3,65,3,113,69,13,1,32,3,65,1,106,33,6,32,3,45,0,0,13,0,12,4,11,11,32,3,65,124,106,33,6,3,64,32,6,65,4,106,34,6,40,2,0,34,3,65,127,115,32,3,65,255,253,251,119,106,113,65,128,129,130,132,120,113,69,13,0,11,32,3,65,255,1,113,69,13,1,3,64,32,6,45,0,1,33,2,32,6,65,1,106,34,3,33,6,32,2,13,0,12,3,11,11,32,0,33,3,12,1,11,32,6,33,3,11,32,3,32,0,107,34,6,65,112,79,13,0,2,64,2,64,2,64,32,6,65,11,79,13,0,32,1,32,6,65,1,116,58,0,48,32,1,65,48,106,65,1,114,34,2,33,3,32,6,13,1,12,2,11,32,6,65,16,106,65,112,113,34,2,16,13,33,3,32,1,32,2,65,1,114,54,2,48,32,1,32,3,54,2,56,32,1,32,6,54,2,52,32,1,65,48,106,65,1,114,33,2,11,32,3,32,0,32,6,16,2,26,11,32,3,32,6,106,65,0,58,0,0,32,1,40,2,56,32,2,32,1,45,0,48,34,3,65,1,113,34,0,27,33,6,2,64,2,64,2,64,2,64,2,64,32,1,40,2,52,32,3,65,1,118,32,0,27,34,3,65,1,71,13,0,32,6,44,0,0,34,5,65,0,72,13,1,65,1,32,1,40,2,24,34,2,107,34,3,65,127,76,13,9,65,255,255,255,255,7,33,4,2,64,65,0,32,2,107,34,0,65,254,255,255,255,3,75,13,0,32,3,32,0,65,1,116,34,4,32,4,32,3,73,27,34,4,69,13,5,11,32,4,16,13,33,3,32,6,45,0,0,33,5,12,11,11,32,3,65,56,79,13,1,11,65,1,32,1,40,2,24,34,2,107,34,4,65,127,76,13,7,32,3,65,128,127,115,33,7,65,255,255,255,255,7,33,0,2,64,65,0,32,2,107,34,5,65,254,255,255,255,3,75,13,0,32,4,32,5,65,1,116,34,0,32,0,32,4,73,27,34,0,69,13,2,11,32,0,16,13,33,4,12,5,11,65,183,1,33,2,32,3,33,0,3,64,32,2,65,1,106,33,2,32,0,65,8,118,34,0,13,0,11,32,2,65,128,2,78,13,3,65,1,16,13,34,0,32,2,58,0,0,32,1,65,24,106,65,8,106,32,0,65,1,106,34,4,54,2,0,32,1,32,0,54,2,24,32,1,32,4,54,2,28,32,1,65,24,106,32,2,65,201,126,106,16,24,32,1,40,2,28,65,127,106,33,2,32,3,33,0,3,64,32,2,32,0,58,0,0,32,2,65,127,106,33,2,32,0,65,8,118,34,0,13,0,11,32,1,40,2,28,33,0,12,5,11,65,0,33,0,65,0,33,4,12,3,11,65,0,33,4,65,0,33,3,12,6,11,16,21,0,11,16,25,16,5,0,11,32,4,32,5,106,34,8,32,7,58,0,0,32,4,32,0,106,33,7,32,8,65,1,106,33,0,2,64,32,5,65,1,72,13,0,32,4,32,2,32,5,16,2,26,32,1,40,2,24,33,2,11,32,1,65,32,106,32,7,54,2,0,32,1,32,0,54,2,28,32,1,32,4,54,2,24,32,2,69,13,0,32,2,16,15,11,32,3,65,1,72,13,3,2,64,2,64,32,3,32,1,65,32,106,40,2,0,34,4,32,0,107,76,13,0,32,0,32,1,40,2,24,34,7,107,34,5,32,3,106,34,8,65,127,76,13,2,32,6,32,3,106,33,9,65,255,255,255,255,7,33,2,2,64,32,4,32,7,107,34,4,65,254,255,255,255,3,75,13,0,32,8,32,4,65,1,116,34,2,32,2,32,8,73,27,34,2,69,13,2,11,32,2,16,13,33,4,12,3,11,3,64,32,0,32,6,45,0,0,58,0,0,32,1,32,1,40,2,28,65,1,106,34,0,54,2,28,32,6,65,1,106,33,6,32,3,65,127,106,34,3,13,0,12,5,11,11,65,0,33,2,65,0,33,4,12,1,11,16,22,0,11,32,4,32,2,106,33,7,32,9,32,5,32,6,107,106,33,8,32,4,32,5,106,34,5,33,2,3,64,32,2,32,6,45,0,0,58,0,0,32,6,65,1,106,33,6,32,2,65,1,106,33,2,32,3,65,127,106,34,3,13,0,11,32,5,32,0,32,1,40,2,24,34,3,107,34,6,107,33,2,2,64,32,6,65,1,72,13,0,32,2,32,3,32,6,16,2,26,11,32,4,32,8,106,33,6,2,64,32,1,40,2,28,32,0,107,34,4,65,1,72,13,0,32,6,32,0,32,4,16,2,26,32,6,32,4,106,33,6,11,32,1,65,32,106,32,7,54,2,0,32,1,32,6,54,2,28,32,1,32,2,54,2,24,32,3,69,13,1,32,3,16,15,12,1,11,32,3,32,0,106,34,6,32,5,58,0,0,32,3,32,4,106,33,4,32,6,65,1,106,33,6,2,64,32,0,65,1,72,13,0,32,3,32,2,32,0,16,2,26,11,32,1,65,32,106,32,4,54,2,0,32,1,32,6,54,2,28,32,1,32,3,54,2,24,32,2,69,13,0,32,2,16,15,11,2,64,2,64,2,64,2,64,32,1,65,40,106,34,8,40,2,0,34,6,32,1,65,36,106,34,9,40,2,0,34,0,70,13,0,3,64,32,6,65,120,106,34,3,40,2,0,34,2,69,13,2,32,3,32,2,65,127,106,34,2,54,2,0,32,2,13,1,32,8,32,3,54,2,0,65,0,33,3,2,64,32,1,40,2,28,32,1,40,2,24,34,4,107,34,0,32,6,65,124,106,40,2,0,34,5,107,34,6,69,13,0,32,6,33,2,3,64,32,3,65,1,106,33,3,32,2,65,8,118,34,2,13,0,11,11,2,64,2,64,32,0,65,1,32,3,65,1,106,32,6,65,56,73,27,34,7,32,0,106,34,2,79,13,0,32,1,65,24,106,32,7,16,24,32,1,40,2,24,33,4,12,1,11,32,0,32,2,77,13,0,32,1,32,4,32,2,106,54,2,28,11,32,4,32,5,106,34,2,32,7,106,32,2,32,6,16,1,26,2,64,2,64,32,6,65,55,75,13,0,32,2,32,6,65,64,106,58,0,0,12,1,11,32,3,65,247,1,106,34,0,65,255,1,75,13,4,32,2,32,0,58,0,0,32,1,40,2,24,32,3,32,5,106,106,33,3,3,64,32,3,32,6,58,0,0,32,3,65,127,106,33,3,32,6,65,8,118,34,6,13,0,11,11,32,8,40,2,0,34,6,32,9,40,2,0,34,0,71,13,0,11,11,2,64,32,1,45,0,48,65,1,113,69,13,0,32,1,65,56,106,40,2,0,16,15,11,32,0,32,6,71,13,2,32,1,40,2,16,32,1,65,8,106,65,1,114,32,1,45,0,8,34,6,65,1,113,34,3,27,32,1,40,2,12,32,6,65,1,118,32,3,27,32,1,40,2,24,34,6,32,1,40,2,28,32,6,107,16,8,2,64,32,0,69,13,0,32,1,65,40,106,32,0,54,2,0,32,0,16,15,11,2,64,32,6,69,13,0,32,1,32,6,54,2,28,32,6,16,15,11,2,64,32,1,45,0,8,65,1,113,69,13,0,32,1,65,16,106,40,2,0,16,15,11,32,1,65,192,0,106,36,0,15,11,16,26,16,5,0,11,16,27,16,5,0,11,16,28,16,5,0,11,188,2,1,6,127,2,64,2,64,2,64,2,64,2,64,32,0,40,2,8,34,2,32,0,40,2,4,34,3,107,32,1,79,13,0,32,3,32,0,40,2,0,34,4,107,34,5,32,1,106,34,6,65,127,76,13,2,65,255,255,255,255,7,33,7,2,64,32,2,32,4,107,34,2,65,254,255,255,255,3,75,13,0,32,6,32,2,65,1,116,34,2,32,2,32,6,73,27,34,7,69,13,2,11,32,7,16,13,33,2,12,3,11,32,0,65,4,106,33,0,3,64,32,3,65,0,58,0,0,32,0,32,0,40,2,0,65,1,106,34,3,54,2,0,32,1,65,127,106,34,1,13,0,12,4,11,11,65,0,33,7,65,0,33,2,12,1,11,16,22,0,11,32,3,32,1,106,33,6,32,2,32,5,106,34,5,33,3,3,64,32,3,65,0,58,0,0,32,3,65,1,106,33,3,32,1,65,127,106,34,1,13,0,11,32,2,32,7,106,33,7,32,2,32,6,32,4,107,106,33,4,32,5,32,0,65,4,106,34,6,40,2,0,32,0,40,2,0,34,1,107,34,3,107,33,2,2,64,32,3,65,1,72,13,0,32,2,32,1,32,3,16,2,26,32,0,40,2,0,33,1,11,32,0,32,2,54,2,0,32,6,32,4,54,2,0,32,0,65,8,106,32,7,54,2,0,32,1,69,13,0,32,1,16,15,15,11,11,41,1,1,127,35,0,65,16,107,34,0,36,0,65,249,8,16,9,32,0,65,10,58,0,15,32,0,65,15,106,65,1,16,10,32,0,65,16,106,36,0,11,41,1,1,127,35,0,65,16,107,34,0,36,0,65,201,8,16,9,32,0,65,10,58,0,15,32,0,65,15,106,65,1,16,10,32,0,65,16,106,36,0,11,41,1,1,127,35,0,65,16,107,34,0,36,0,65,221,8,16,9,32,0,65,10,58,0,15,32,0,65,15,106,65,1,16,10,32,0,65,16,106,36,0,11,41,1,1,127,35,0,65,16,107,34,0,36,0,65,178,8,16,9,32,0,65,10,58,0,15,32,0,65,15,106,65,1,16,10,32,0,65,16,106,36,0,11,149,17,2,35,127,5,126,35,0,65,16,107,34,3,36,0,32,0,32,1,41,3,0,55,3,0,32,0,32,1,41,3,32,55,3,32,32,0,65,24,106,32,1,65,24,106,41,3,0,55,3,0,32,0,65,16,106,32,1,65,16,106,41,3,0,55,3,0,32,0,65,8,106,32,1,65,8,106,41,3,0,55,3,0,2,64,32,2,40,2,0,34,1,65,127,74,13,0,32,3,65,140,8,54,2,8,65,131,10,65,3,113,33,4,65,130,10,65,3,113,33,5,65,132,10,65,3,113,33,6,65,133,10,65,3,113,33,7,65,134,10,65,3,113,33,8,65,135,10,65,3,113,33,9,65,136,10,65,3,113,33,10,65,137,10,65,3,113,33,11,65,138,10,65,3,113,33,12,65,139,10,65,3,113,33,13,65,140,10,65,3,113,33,14,65,141,10,65,3,113,33,15,65,142,10,65,3,113,33,16,65,143,10,65,3,113,33,17,65,144,10,65,3,113,33,18,65,145,10,65,3,113,33,19,65,146,10,65,3,113,33,20,65,147,10,65,3,113,33,21,65,148,10,65,3,113,33,22,65,149,10,65,3,113,33,23,65,150,10,65,3,113,33,24,65,151,10,65,3,113,33,25,65,152,10,65,3,113,33,26,65,153,10,65,3,113,33,27,65,154,10,65,3,113,33,28,65,155,10,65,3,113,33,29,65,156,10,65,3,113,33,30,65,157,10,65,3,113,33,31,65,158,10,65,3,113,33,32,65,159,10,65,3,113,33,33,65,160,10,65,3,113,33,34,65,161,10,65,3,113,33,35,65,162,10,65,3,113,33,36,65,163,10,33,1,2,64,2,64,65,129,10,65,3,113,34,37,69,13,0,32,5,65,0,71,32,4,65,0,71,113,32,6,65,0,71,113,32,7,65,0,71,113,32,8,65,0,71,113,32,9,65,0,71,113,32,10,65,0,71,113,32,11,65,0,71,113,32,12,65,0,71,113,32,13,65,0,71,113,32,14,65,0,71,113,32,15,65,0,71,113,32,16,65,0,71,113,32,17,65,0,71,113,32,18,65,0,71,113,32,19,65,0,71,113,32,20,65,0,71,113,32,21,65,0,71,113,32,22,65,0,71,113,32,23,65,0,71,113,32,24,65,0,71,113,32,25,65,0,71,113,32,26,65,0,71,113,32,27,65,0,71,113,32,28,65,0,71,113,32,29,65,0,71,113,32,30,65,0,71,113,32,31,65,0,71,113,32,32,65,0,71,113,32,33,65,0,71,113,32,34,65,0,71,113,32,35,65,0,71,113,32,36,65,0,71,113,65,163,10,65,3,113,65,0,71,113,13,1,11,65,131,10,65,130,10,32,5,27,65,132,10,32,5,69,32,4,69,114,34,1,27,65,133,10,32,1,32,6,69,114,34,1,27,65,134,10,32,1,32,7,69,114,34,1,27,65,135,10,32,1,32,8,69,114,34,1,27,65,136,10,32,1,32,9,69,114,34,1,27,65,137,10,32,1,32,10,69,114,34,1,65,1,113,27,65,138,10,32,1,32,11,69,114,34,1,65,1,113,27,65,139,10,32,1,32,12,69,114,34,1,65,1,113,27,65,140,10,32,1,32,13,69,114,34,1,65,1,113,27,65,141,10,32,1,32,14,69,114,34,1,65,1,113,27,65,142,10,32,1,32,15,69,114,34,1,65,1,113,27,65,143,10,32,1,32,16,69,114,34,1,65,1,113,27,65,144,10,32,1,32,17,69,114,34,1,65,1,113,27,65,145,10,32,1,32,18,69,114,34,1,65,1,113,27,65,146,10,32,1,32,19,69,114,34,1,65,1,113,27,65,147,10,32,1,32,20,69,114,34,1,65,1,113,27,65,148,10,32,1,32,21,69,114,34,1,65,1,113,27,65,149,10,32,1,32,22,69,114,34,1,65,1,113,27,65,150,10,32,1,32,23,69,114,34,1,65,1,113,27,65,151,10,32,1,32,24,69,114,34,1,65,1,113,27,65,152,10,32,1,32,25,69,114,34,1,65,1,113,27,65,153,10,32,1,32,26,69,114,34,1,65,1,113,27,65,154,10,32,1,32,27,69,114,34,1,65,1,113,27,65,155,10,32,1,32,28,69,114,34,1,65,1,113,27,65,156,10,32,1,32,29,69,114,34,1,65,1,113,27,65,157,10,32,1,32,30,69,114,34,1,65,1,113,27,65,158,10,32,1,32,31,69,114,34,1,65,1,113,27,65,159,10,32,1,32,32,69,114,34,1,65,1,113,27,65,160,10,32,1,32,33,69,114,34,1,65,1,113,27,65,161,10,32,1,32,34,69,114,34,1,65,1,113,27,65,162,10,32,1,32,35,69,114,34,1,65,1,113,27,65,163,10,32,1,32,36,69,114,65,1,113,27,65,129,10,32,37,27,65,124,106,33,1,3,64,32,1,65,4,106,34,1,40,2,0,34,5,65,127,115,32,5,65,255,253,251,119,106,113,65,128,129,130,132,120,113,69,13,0,11,32,5,65,255,1,113,69,13,0,32,1,33,5,3,64,32,5,45,0,1,33,4,32,5,65,1,106,34,1,33,5,32,4,13,0,11,11,32,1,65,129,10,107,34,1,65,13,106,16,13,34,5,32,1,54,2,4,32,5,32,1,54,2,0,32,5,65,0,54,2,8,32,5,65,12,106,34,5,65,129,10,32,1,65,1,106,16,2,26,32,3,65,160,8,54,2,8,32,3,32,5,54,2,12,32,3,65,8,106,16,3,32,3,65,140,8,54,2,8,32,3,40,2,12,34,5,65,124,106,34,1,32,1,40,2,0,65,127,106,34,1,54,2,0,2,64,32,1,65,127,74,13,0,32,5,65,116,106,16,15,11,32,2,40,2,0,33,1,11,2,64,32,1,69,13,0,2,64,2,64,2,64,2,64,2,64,2,64,32,1,173,66,7,131,66,0,81,13,0,2,64,32,0,65,32,106,41,3,0,34,38,167,34,5,65,1,71,13,0,32,0,41,3,0,80,13,6,11,32,1,172,34,39,66,6,136,32,39,66,63,135,66,58,134,132,33,40,32,1,65,63,113,34,4,173,33,41,32,5,33,1,2,64,32,4,69,13,0,32,0,32,5,65,3,116,106,65,120,106,41,3,0,66,192,0,32,41,125,136,66,0,82,32,5,106,33,1,11,32,0,65,32,106,34,4,32,1,32,40,167,34,8,106,34,6,65,4,32,6,65,4,73,27,34,1,173,55,3,0,32,40,32,6,173,34,39,88,13,1,32,0,66,0,55,3,0,32,4,66,1,55,3,0,12,5,11,2,64,32,0,65,32,106,40,2,0,34,5,65,1,71,13,0,32,0,41,3,0,80,13,5,11,32,1,65,6,117,33,2,32,5,33,4,2,64,32,1,65,63,113,34,6,69,13,0,32,0,32,5,65,3,116,106,65,120,106,41,3,0,66,192,0,32,6,173,125,136,66,0,82,32,5,106,33,4,11,32,0,65,32,106,32,4,32,2,106,34,4,65,4,32,4,65,4,73,27,34,4,173,55,3,0,2,64,32,4,32,5,70,13,0,32,0,32,4,65,3,116,106,65,120,106,66,0,55,3,0,11,32,4,65,3,116,34,4,32,1,65,3,117,34,1,77,13,1,32,0,32,1,106,32,0,32,4,32,1,107,34,4,32,5,65,3,116,34,5,32,4,32,5,73,27,16,1,26,32,0,65,0,32,1,16,4,26,12,4,11,32,6,32,1,107,33,2,32,6,65,4,75,13,2,32,0,32,2,65,127,115,34,1,32,5,106,65,3,116,106,41,3,0,33,42,32,38,66,255,255,255,255,15,131,32,40,124,32,39,90,13,1,32,0,32,6,32,1,106,65,3,116,106,32,42,66,192,0,32,41,125,136,55,3,0,32,6,65,127,106,33,6,12,2,11,32,0,66,0,55,3,0,32,0,65,32,106,66,1,55,3,0,12,2,11,32,0,32,6,32,1,106,65,3,116,106,34,1,32,42,32,41,134,34,39,55,3,0,2,64,32,5,65,2,73,13,0,32,1,32,0,32,5,65,126,106,32,2,107,65,3,116,106,41,3,0,66,192,0,32,41,125,136,32,39,132,55,3,0,11,32,2,65,1,106,33,2,11,2,64,32,40,66,2,124,34,38,32,6,32,2,107,173,34,39,86,13,0,32,0,32,6,65,126,106,32,2,32,8,106,107,65,3,116,106,33,1,65,0,32,2,107,33,4,66,192,0,32,41,125,33,42,32,0,32,6,65,127,106,32,2,107,65,3,116,106,33,5,3,64,32,5,32,1,65,8,106,41,3,0,32,41,134,34,39,55,3,0,32,5,32,1,41,3,0,32,42,136,32,39,132,55,3,0,32,5,65,120,106,33,5,32,1,65,120,106,33,1,32,6,32,4,106,33,2,32,4,65,127,106,34,7,33,4,32,38,32,2,65,127,106,173,34,39,88,13,0,11,65,0,32,7,107,33,2,11,2,64,32,40,66,1,124,32,39,86,13,0,32,0,32,6,32,2,65,127,115,106,34,1,65,3,116,106,32,0,32,1,32,8,107,65,3,116,106,41,3,0,32,41,134,55,3,0,32,2,65,1,106,33,2,11,32,2,32,6,79,13,0,32,6,32,2,107,33,5,32,0,32,6,65,127,106,32,2,107,65,3,116,106,33,1,3,64,32,1,66,0,55,3,0,32,1,65,120,106,33,1,32,5,65,127,106,34,5,13,0,11,11,32,0,65,32,106,34,5,41,3,0,34,41,66,127,124,34,39,80,13,0,32,0,32,41,167,65,3,116,106,65,120,106,33,1,3,64,32,1,41,3,0,66,0,82,13,1,32,5,32,39,55,3,0,32,1,65,120,106,33,1,32,39,66,127,124,34,39,66,0,82,13,0,11,11,32,3,65,16,106,36,0,11,180,5,2,3,127,2,126,2,64,2,64,2,64,32,0,32,2,70,13,0,32,0,32,1,70,13,1,32,0,32,1,41,3,32,55,3,32,32,0,32,1,32,1,40,2,32,65,3,116,16,2,26,32,0,32,2,40,2,32,34,3,32,0,40,2,32,34,4,32,4,32,3,73,27,34,5,65,4,32,5,65,4,73,27,173,55,3,32,2,64,32,4,32,3,79,13,0,32,0,32,4,65,3,116,106,33,1,3,64,32,1,66,0,55,3,0,32,1,65,8,106,33,1,32,4,65,1,106,34,4,32,5,73,13,0,11,11,2,64,32,3,69,13,0,32,0,33,1,3,64,32,1,32,2,41,3,0,32,1,41,3,0,132,55,3,0,32,1,65,8,106,33,1,32,2,65,8,106,33,2,32,3,65,127,106,34,3,13,0,11,11,32,0,65,32,106,34,1,41,3,0,34,6,66,127,124,34,7,80,13,2,32,0,32,6,167,65,3,116,106,65,120,106,33,2,3,64,32,2,41,3,0,66,0,82,13,3,32,1,32,7,55,3,0,32,2,65,120,106,33,2,32,7,66,127,124,34,7,66,0,82,13,0,12,3,11,11,32,0,32,1,40,2,32,34,3,32,0,40,2,32,34,4,32,4,32,3,73,27,34,5,65,4,32,5,65,4,73,27,173,55,3,32,2,64,32,4,32,3,79,13,0,32,0,32,4,65,3,116,106,33,2,3,64,32,2,66,0,55,3,0,32,2,65,8,106,33,2,32,4,65,1,106,34,4,32,5,73,13,0,11,11,2,64,32,3,69,13,0,32,0,33,2,3,64,32,2,32,1,41,3,0,32,2,41,3,0,132,55,3,0,32,2,65,8,106,33,2,32,1,65,8,106,33,1,32,3,65,127,106,34,3,13,0,11,11,32,0,65,32,106,34,1,41,3,0,34,6,66,127,124,34,7,80,13,1,32,0,32,6,167,65,3,116,106,65,120,106,33,2,3,64,32,2,41,3,0,66,0,82,13,2,32,1,32,7,55,3,0,32,2,65,120,106,33,2,32,7,66,127,124,34,7,80,69,13,0,12,2,11,11,32,0,32,2,40,2,32,34,3,32,0,40,2,32,34,4,32,4,32,3,73,27,34,5,65,4,32,5,65,4,73,27,173,55,3,32,2,64,32,4,32,3,79,13,0,32,0,32,4,65,3,116,106,33,1,3,64,32,1,66,0,55,3,0,32,1,65,8,106,33,1,32,4,65,1,106,34,4,32,5,73,13,0,11,11,2,64,32,3,69,13,0,32,0,33,1,3,64,32,1,32,2,41,3,0,32,1,41,3,0,132,55,3,0,32,1,65,8,106,33,1,32,2,65,8,106,33,2,32,3,65,127,106,34,3,13,0,11,11,32,0,65,32,106,34,1,41,3,0,34,6,66,127,124,34,7,80,13,0,32,0,32,6,167,65,3,116,106,65,120,106,33,2,3,64,32,2,41,3,0,66,0,82,13,1,32,1,32,7,55,3,0,32,2,65,120,106,33,2,32,7,66,127,124,34,7,80,69,13,0,11,11,11,196,1,1,1,127,35,0,65,16,107,34,4,36,0,65,221,9,16,9,32,4,65,32,58,0,15,32,4,65,15,106,65,1,16,10,32,0,40,2,0,16,9,32,4,65,32,58,0,10,32,4,65,10,106,65,1,16,10,65,239,9,16,9,32,4,65,32,58,0,9,32,4,65,9,106,65,1,16,10,32,1,40,2,0,16,9,32,4,65,32,58,0,11,32,4,65,11,106,65,1,16,10,65,245,9,16,9,32,4,65,32,58,0,14,32,4,65,14,106,65,1,16,10,32,2,53,2,0,16,11,32,4,65,32,58,0,13,32,4,65,13,106,65,1,16,10,65,251,9,16,9,32,4,65,32,58,0,12,32,4,65,12,106,65,1,16,10,32,3,40,2,0,16,9,32,4,65,10,58,0,8,32,4,65,8,106,65,1,16,10,32,4,65,16,106,36,0,11,210,2,3,2,127,1,126,1,127,35,0,65,240,1,107,34,0,36,0,32,0,65,48,106,16,0,32,0,66,1,55,3,32,32,0,66,0,55,3,0,65,0,33,1,3,64,32,0,65,48,106,32,1,106,49,0,0,33,2,32,0,65,8,54,2,92,32,0,65,224,0,106,32,0,32,0,65,220,0,106,16,29,32,0,65,144,1,106,65,32,106,34,3,66,1,55,3,0,32,0,66,0,55,3,144,1,32,0,65,192,1,106,65,32,106,66,1,55,3,0,32,0,32,2,55,3,192,1,32,0,65,144,1,106,32,0,65,224,0,106,32,0,65,192,1,106,16,30,32,0,65,32,106,32,3,41,3,0,55,3,0,32,0,65,24,106,32,0,65,144,1,106,65,24,106,41,3,0,55,3,0,32,0,65,16,106,32,0,65,144,1,106,65,16,106,41,3,0,55,3,0,32,0,32,0,41,3,152,1,55,3,8,32,0,32,0,41,3,144,1,55,3,0,32,1,65,1,106,34,1,65,32,71,13,0,11,32,0,65,32,106,40,2,0,33,1,32,0,41,3,0,33,2,32,0,65,4,54,2,144,1,32,0,65,145,9,54,2,192,1,32,0,65,156,9,54,2,96,32,0,65,216,9,54,2,48,2,64,32,1,65,1,71,13,0,32,2,80,69,13,0,32,0,65,240,1,106,36,0,15,11,32,0,65,192,1,106,32,0,65,48,106,32,0,65,144,1,106,32,0,65,224,0,106,16,31,16,5,0,11,216,2,3,2,127,1,126,1,127,35,0,65,240,1,107,34,1,36,0,32,1,65,48,106,16,0,32,1,66,1,55,3,32,32,1,66,0,55,3,0,65,0,33,2,3,64,32,1,65,48,106,32,2,106,49,0,0,33,3,32,1,65,8,54,2,92,32,1,65,224,0,106,32,1,32,1,65,220,0,106,16,29,32,1,65,144,1,106,65,32,106,34,4,66,1,55,3,0,32,1,66,0,55,3,144,1,32,1,65,192,1,106,65,32,106,66,1,55,3,0,32,1,32,3,55,3,192,1,32,1,65,144,1,106,32,1,65,224,0,106,32,1,65,192,1,106,16,30,32,1,65,32,106,32,4,41,3,0,55,3,0,32,1,65,24,106,32,1,65,144,1,106,65,24,106,41,3,0,55,3,0,32,1,65,16,106,32,1,65,144,1,106,65,16,106,41,3,0,55,3,0,32,1,32,1,41,3,152,1,55,3,8,32,1,32,1,41,3,144,1,55,3,0,32,2,65,1,106,34,2,65,32,71,13,0,11,32,1,65,32,106,40,2,0,33,2,32,1,41,3,0,33,3,32,1,65,9,54,2,144,1,32,1,65,145,9,54,2,192,1,32,1,65,156,9,54,2,96,32,1,65,164,10,54,2,48,2,64,32,2,65,1,71,13,0,32,3,80,69,13,0,32,0,16,23,32,1,65,240,1,106,36,0,32,0,15,11,32,1,65,192,1,106,32,1,65,48,106,32,1,65,144,1,106,32,1,65,224,0,106,16,31,16,5,0,11,216,2,3,2,127,1,126,1,127,35,0,65,240,1,107,34,1,36,0,32,1,65,48,106,16,0,32,1,66,1,55,3,32,32,1,66,0,55,3,0,65,0,33,2,3,64,32,1,65,48,106,32,2,106,49,0,0,33,3,32,1,65,8,54,2,92,32,1,65,224,0,106,32,1,32,1,65,220,0,106,16,29,32,1,65,144,1,106,65,32,106,34,4,66,1,55,3,0,32,1,66,0,55,3,144,1,32,1,65,192,1,106,65,32,106,66,1,55,3,0,32,1,32,3,55,3,192,1,32,1,65,144,1,106,32,1,65,224,0,106,32,1,65,192,1,106,16,30,32,1,65,32,106,32,4,41,3,0,55,3,0,32,1,65,24,106,32,1,65,144,1,106,65,24,106,41,3,0,55,3,0,32,1,65,16,106,32,1,65,144,1,106,65,16,106,41,3,0,55,3,0,32,1,32,1,41,3,152,1,55,3,8,32,1,32,1,41,3,144,1,55,3,0,32,2,65,1,106,34,2,65,32,71,13,0,11,32,1,65,32,106,40,2,0,33,2,32,1,41,3,0,33,3,32,1,65,16,54,2,144,1,32,1,65,145,9,54,2,192,1,32,1,65,156,9,54,2,96,32,1,65,177,10,54,2,48,2,64,32,2,65,1,71,13,0,32,3,80,69,13,0,32,0,16,23,32,1,65,240,1,106,36,0,66,50,15,11,32,1,65,192,1,106,32,1,65,48,106,32,1,65,144,1,106,32,1,65,224,0,106,16,31,16,5,0,11,216,2,3,2,127,1,126,1,127,35,0,65,240,1,107,34,1,36,0,32,1,65,48,106,16,0,32,1,66,1,55,3,32,32,1,66,0,55,3,0,65,0,33,2,3,64,32,1,65,48,106,32,2,106,49,0,0,33,3,32,1,65,8,54,2,92,32,1,65,224,0,106,32,1,32,1,65,220,0,106,16,29,32,1,65,144,1,106,65,32,106,34,4,66,1,55,3,0,32,1,66,0,55,3,144,1,32,1,65,192,1,106,65,32,106,66,1,55,3,0,32,1,32,3,55,3,192,1,32,1,65,144,1,106,32,1,65,224,0,106,32,1,65,192,1,106,16,30,32,1,65,32,106,32,4,41,3,0,55,3,0,32,1,65,24,106,32,1,65,144,1,106,65,24,106,41,3,0,55,3,0,32,1,65,16,106,32,1,65,144,1,106,65,16,106,41,3,0,55,3,0,32,1,32,1,41,3,152,1,55,3,8,32,1,32,1,41,3,144,1,55,3,0,32,2,65,1,106,34,2,65,32,71,13,0,11,32,1,65,32,106,40,2,0,33,2,32,1,41,3,0,33,3,32,1,65,24,54,2,144,1,32,1,65,145,9,54,2,192,1,32,1,65,156,9,54,2,96,32,1,65,187,10,54,2,48,2,64,32,2,65,1,71,13,0,32,3,80,69,13,0,32,0,16,23,32,1,65,240,1,106,36,0,66,50,15,11,32,1,65,192,1,106,32,1,65,48,106,32,1,65,144,1,106,32,1,65,224,0,106,16,31,16,5,0,11,11,179,3,18,0,65,128,8,11,4,0,0,0,0,0,65,132,8,11,40,0,0,0,0,0,0,0,0,1,0,0,0,2,0,0,0,3,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,4,0,0,0,3,0,0,0,0,65,172,8,11,6,116,111,112,105,99,0,0,65,178,8,11,23,108,105,115,116,83,116,97,99,107,32,105,115,32,110,111,116,32,101,109,112,116,121,0,0,65,201,8,11,20,105,116,101,109,67,111,117,110,116,32,116,111,111,32,108,97,114,103,101,0,0,65,221,8,11,28,105,116,101,109,67,111,117,110,116,32,116,111,111,32,108,97,114,103,101,32,102,111,114,32,82,76,80,0,0,65,249,8,11,24,67,111,117,110,116,32,116,111,111,32,108,97,114,103,101,32,102,111,114,32,82,76,80,0,0,65,145,9,11,11,116,101,109,112,118,32,61,61,32,48,0,0,65,156,9,11,60,47,118,97,114,47,102,111,108,100,101,114,115,47,116,116,47,56,118,52,98,98,48,51,120,54,56,53,99,52,57,106,107,95,118,107,118,114,55,115,56,48,48,48,48,103,112,47,84,47,47,101,118,101,110,116,46,99,112,112,0,0,65,216,9,11,5,105,110,105,116,0,0,65,221,9,11,18,65,115,115,101,114,116,105,111,110,32,102,97,105,108,101,100,58,0,0,65,239,9,11,6,102,117,110,99,58,0,0,65,245,9,11,6,108,105,110,101,58,0,0,65,251,9,11,6,102,105,108,101,58,0,0,65,129,10,11,35,67,97,110,32,110,111,116,32,115,104,105,102,116,32,98,121,32,97,32,110,101,103,97,116,105,118,101,32,118,97,108,117,101,46,0,0,65,164,10,11,13,114,101,116,117,114,110,83,116,114,105,110,103,0,0,65,177,10,11,10,114,101,116,117,114,110,73,110,116,0,0,65,187,10,11,11,114,101,116,117,114,110,85,105,110,116,0},
		abi:
`[
    {
        "name": "init",
        "inputs": [],
        "outputs": [],
        "constant": "true",
        "payable": "false",
        "type": "function"
    },
    {
        "name": "returnString",
        "inputs": [
            {
                "name": "name",
                "type": "string"
            }
        ],
        "outputs": [
            {
                "name": "",
                "type": "string"
            }
        ],
        "constant": "true",
        "payable": "false",
        "type": "function"
    },
    {
        "name": "returnInt",
        "inputs": [
            {
                "name": "name",
                "type": "string"
            }
        ],
        "outputs": [
            {
                "name": "",
                "type": "int64"
            }
        ],
        "constant": "true",
        "payable": "false",
        "type": "function"
    },
    {
        "name": "returnUint",
        "inputs": [
            {
                "name": "name",
                "type": "string"
            }
        ],
        "outputs": [
            {
                "name": "",
                "type": "uint64"
            }
        ],
        "constant": "true",
        "payable": "false",
        "type": "function"
    },
    {
        "name": "topic",
        "inputs": [
            {
                "type": "string"
            }
        ],
        "type": "event"
    }
]`,
		hpp:
`#pragma once
#include <dipc/dipc.hpp>
using namespace dipc;

class envEvent : public Contract {
 public: 
  CONSTANT void init();
  CONSTANT char* returnString(char* name);
  CONSTANT int64_t returnInt(char* name);
  CONSTANT uint64_t returnUint(char* name);
};

DIPC_EVENT(topic, const char*);`,
		cpp:
`#include "event.hpp"

CONSTANT void envEvent::init() {}

CONSTANT char* envEvent::returnString(char* name) {
    DIPC_EMIT_EVENT(topic, name);
    return name;
}

CONSTANT int64_t envEvent::returnInt(char* name) {
    DIPC_EMIT_EVENT(topic, name);
    int64_t num = 50;
    return num;
}

CONSTANT uint64_t envEvent::returnUint(char* name) {
    DIPC_EMIT_EVENT(topic, name);
    uint64_t num = 50;
    return num;
}`,
	},

}

func GetTestData(contractName string ) (code []byte, abi []byte)  {
	codeAndAbi := contractTestData[contractName]
	return codeAndAbi.code, []byte(codeAndAbi.abi)
}
