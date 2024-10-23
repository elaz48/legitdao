package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	units     string
	referrals []string
}

type treeUint struct {
	units     uint64
	referred  uint64
	top       uint64
	referrals []string
}

const clevelDivider = 10000
const totalEmissionWithCents = 10000000000000000
const founderShare = 0.15

func main() {
	data := `
Everyone not referred by someone is referred by the team address:
9Fa05D3823Cb5227d75FeC60aF97e0616aCa6367
------------------------------------------------------------
0) 6db0ba283092dbb17e30639193dbf67091d16ad0
units: 14050000
referrals:
----
1) 48e36BdBBc54219136FDd20685ccC73166Fe8324
units: 21360000
referrals: aaf4394afbe1b7c6c8379da08a999da788a9728c
----
2) cfc8be6e4ff4e75ca8c95ed1a7efc4ea8d894d0f
units: 108669000
referrals:
----
3) 41224e44922d2d98aa6360c705772c59aacec96b
units: 47040000
referrals: eec9aed9d7d5430dbee4f6425c5974a6dce5dd78,
a6e9ce90646c5624c7cd530422eaa3a6662f1992
----
4) b4536dc4950b53e3c16fb85c5a56f9be239a7b6f
units: 29600000
referrals:
----
5) 6390b88a6369d08e28ece4d741131d6ffbac224c
units: 5620000
----
6) 8818518f5250ed3dc7bcf4e539f6b35b54e29b29
units: 281000000
referrals: 014b5bb2f3800637e0e7de077a9af4803ab81e9c
----
7) 93967805faa6eabe9de62d578b1f076863908133
units: 267000000
referrals: e53b3cb82815b613062dc657a9fcd6cb77ec595a,
172aa24dcaefcb2ef14104761d37233ca6e6c166,
423bdfeb5d25b0c8d5c6ff598e99eb883a340381
----
8) 844950422e1f7306e536fa68a6eff106151cff24
units: 5340000
referrals:
----
9) 5176fe6b79a153beba945f1ba0fb2fdf4f6a8dee
units: 56200000
referrals:
----
10) 26bc99a54c290a7865dcbf0b1ed8a4f013ea964f
units: 281000000
referrals:
----
11) 53daccec0fda34c30f89cd64ad29a9bbe84d275e
units: 95540000
referrals:
----
12) 5db0ce4be877074f838ec85eb196784ee529df15
units: 256834000
referrals:
----
13) 40cc689036296a1d064805a5b913043a5177bc4e
units: 240000000
referrals:
----
14) 5b5d95d32f55f08480704077888d713d50c8f18f
units: 53400000
referrals:
----
15) 96bad93f26c64e72f7b22e3dee731ca51e84c78b
units: 11099500
referrals:
----
16) 6782f058b8cc66625c9da50137cc71730bde1c08
units: 267000000
referrals:
----
17) 6890eef74ea683a9383b6a32099b3967eae7907e
units: 266199000
referrals: c3646c9f183a7ff16addb0b778221519a1faeac4
----
18) 5159f99d35cb74796de5664da456b79a21ac5449
units: 267000000
referrals:
----
19) 5dce7412dac1860b1c9eaceae3a09ed2f4ec7600
units: 112400000
referrals:
----
20) 8de7aa0a77134b3e5e09f3586ad6769f79ea8b2d
units: 14050000
referrals:
----
21) 503bb794fe7a57a15fb7e7967af22bf341682c5b
units: 266466000
referrals:
----
22) b21d65b8dd5d419592915a1d9bf281b7036e4ca7
units: 59291000
referrals:
----
23) 8c9a61a81863744a3c778f497f2853c9ee2c9d50
units: 81490000
referrals:
----
24) ca9e4ec800b85345b186dcda5c02671a861dc067
units: 146583000
referrals: 32f0281e2ab2c728ac592c5006b5cdf35a80b0e3,
43ddc99e9baf04b12979b1ecc9272a0585b27c90,
f81526ac03a5f38ff58eb9cd4b5caaa3d8ee6771,
880a9d891b285650defe2969ed22608d26ddfff3,
8e367218246f9df464dd925169c470d8f0622222,
49ec0ed807417b69e2950fa809a60e1994d2dc6b,
69e1181e26fab42ab0f73797aa8e8b8dc77628cf
----
25) 7245aa3377401c1a0482219053edc8545d65920f
units: 148000000
referrals:
----
26) 1689ab3994e413046f3d8e153768c7757a015420
units: 56200000
referrals:
----
27) 337cf9b5bb45285eec16661d120c8ab51134c9aa
units: 264330000
referrals: 6789297f7623d5cd5a274df4e82dc0ddc0b96948,
44ee1512f5d7455523cda8338b04ae6a60a99de4
----
28) 43234e1a1c49023725af693b6bc5a76c11c46765
units: 26700000
referrals:
----
29) c9e24327b2bd18dd62d870a993a273c16c895cf6
units: 118400000
referrals:
----
30) ed116130ec54b9be4d58ac17a3849e3df497125a
units: 281000000
referrals:
----
31) 35867f6474a4781ecb49d383acdf6b929af143dd
units: 227217000
referrals: 3e091d480d0e3429fab8e072075b7502f218ebb8
----
32) ebacba195accc94e6ac9b34207f3916c41119329
units: 267000000
referrals: 337cf9b5bb45285eec16661d120c8ab51134c9aa
----
33) 89cbf332571f54e2d1ceaddc2b91b6a96f21df00
units: 74000000
referrals: 0730509fac1c163e947c26cfa8a40fe76cc846d4
----
34) 61b0660f216ff73f069f4f68ec5e8ab8795dfe7e
units: 14800000
referrals:
----
35) 32f0281e2ab2c728ac592c5006b5cdf35a80b0e3
units: 245400000
referrals:
----
36) 621cb7f95cd54b6884ca331cfb5e610fa5cb092d
units: 245373000
referrals:
----
37) 3e2ea46bdfe5ec7f1a12b3f24d47439e1b5e9ebf
units: 296000000
referrals:
----
38) 128c688929b28ffd468705220ecdd71a74e81b89
units: 28100000
referrals:
----
39) c671e68226ae0aa544b6df8f2f16cf6232e187ea
units: 140175000
referrals:
----
40) 546b76e990ad05f904c7e1ef648cab8234c7311c
units: 48000000
referrals:
----
41) e443ee342ccd35b2ba7018aa9816abeaa8123586
units: 26400000
referrals:
----
42) e8a1a820cdc94ca5e2726612a8dc0e7b0e33adaa
units: 296000000
referrals:
----
43) 23c8dc48eb57f7748d2ed89886581ff95978efd6
units: 24000000
referrals:
----
44) a3ec1ac881f6397d7ac0916c91013a159b2fa9a7
units: 59200000
referrals: 61b0660f216ff73f069f4f68ec5e8ab8795dfe7e
----
45) 093ff2fb6b3c815ff955248519ec7b7af42e665e
units: 2810000
referrals: bb47df9c00cfc201e709764f31306d02220aa1b4
----
46) 941bbd3d449cd092781f6f12819dabd537c36947
units: 59010000
referrals:
----
47) 1e8c82755f62081cc7c760730677c5776da27c32
units: 281000000
referrals:
----
48) b07a67678bab935c2ce60edc41463aa5785e9dd4
units: 43555000
referrals:
----
49) 0dab09ffe73b6282d8b38479d6fc0472ed16118a
units: 267000000
referrals: 61229a2ac51b2430d228f58a2658e1770346f0f6,
c531e47f882a34543cc77e92ffbba93d3ea11117,
b2e9113d7ac34f7d39662b0e8ddc769bdec9cbf7,
3023138cb3aeeae67c4033015a39b56af582e3fc,
7be8c96a98cb2fe2fd6c59edd2a83cebb3233ba2,
871d583deb3083993b28487beb934eef155e0118
----
50) 178132f6c553ce96d4b47ec224a5c37e76be39dd
units: 240000000
referrals:
----
51) 058f3f556ce5e6364dee07b85087de0896251544
units: 240000000
referrals:
----
52) d9870a9677778db6da1da141d9faf3cc8c1dbe15
units: 176100000
referrals: dc9e7a6a086dbb784c7bd7d84b4884b0ea1f168a
----
53) 8dd3382eb29c4eccf48482d7d7ed93301638b702
units: 9176000
referrals:
----
54) c4ce1c95f0ed3ae7f1428712c1e768cef2462ec2
units: 296000000
referrals:
----
55) 61178ed786321f6acdf01633e336d94fc42ec53a
units: 26700000
referrals:
----
56) 78da7fca531d4b6710d616cdd63d77bc80e7fd32
units: 296000000
referrals: fad9ca41589b3657d56e1a6b38601dcbc5dab508,
8d24e10d251a575cebe15a18c8f2e9e956f57b60,
04559347c9c78805cb65f6b290c0f6fb9fd2c3fd,
dcf8023f18409005207026de7f0d1e780a94f742,
ea24c5b3b0f9c0b5bb93d6a72747c385a87aca7b,
08fc58b846e69b332d6de7facba6062648490e8d,
91c62ff6e40b9718af781148fac9d3b6684a4c24
----
57) ea24c5b3b0f9c0b5bb93d6a72747c385a87aca7b
units: 79032000
referrals:
----
58) ec931d7bef828d5beadac291490a71fae80a401b
units: 55919000
referrals:
----
59) d34a20e69403aa10913dd91455aee574f7611527
units: 110070000
referrals:
----
60) e7fd8119ff58dc422da6c6ecfd0f8a16b4610062
units: 250416000
referrals:
----
61) ef4af552b493cd52ca79765ac29098e5b65c3989
units: 281000000
referrals: 25e7d4c6529dc774ae1a44e686084273dab679fb,
8c9a61a81863744a3c778f497f2853c9ee2c9d50,
35867f6474a4781ecb49d383acdf6b929af143dd,
cfc8be6e4ff4e75ca8c95ed1a7efc4ea8d894d0f,
475131c7992a6666969b344ad8cb8b9f8e98e606,
c555acf73a1164f7f2eee5e7f59fee502f17afec
----
62) 1c2Cc5C54F3362Ff2bFDc03c0B1f9664b2073563
units: 261664000
referrals:
----
63) 9944aC84145836C679acA9968992E12A9330Ed2c
units: 9835000
referrals:
----
64) b2e9113d7ac34f7d39662b0e8ddc769bdec9cbf7
units: 47793000
referrals:
----
65) 79d4d61060c4071cc8530b47ef7b93209a6d0c33
units: 248310000
referrals:
----
66) b3cd225e39779c41dc994eb881b696f0cbc68e08
units: 240000000
referrals:
----
67) 2ea252ba869eed2a248240a00207583f934404ad
units: 152440000
referrals:
----
68) deda2f75187f96965eb24fd53123f42a5ef342a8
units: 88800000
referrals:
----
69) adc0b7a58ba5fbd9e8a6980350b39c52704a2f6a
units: 126450000
referrals: 8de7aa0a77134b3e5e09f3586ad6769f79ea8b2d
----
70) d0a26322abfaa561da11713e76af4e5611b9a4fd
units: 31506000
referrals:
----
71) 899530f3464a37920674b536384c299c6a09eeee
units: 46992000
referrals:
----
72) 43ddc99e9baf04b12979b1ecc9272a0585b27c90
units: 168210000
referrals:
----
73) b7d0fa0669730e46f9fae6d3a89bc26832b55dfe
units: 89040000
referrals:
----
74) 4ef240f6c76f73b1ac2d068f0309987a9f98a9e3
units: 296000000
referrals:
----
75) ea9d1d1fa0b7bed82f0d2723883f0543a8a18e9a
units: 57720000
referrals:
----
76) 2f5b28dd625ca194894165edaa4704ee75d4e4ab
units: 279424000
referrals: 0dab09ffe73b6282d8b38479d6fc0472ed16118a,
3603b37948c74b6e15c7988aae9624741e7b2398,
c671e68226ae0aa544b6df8f2f16cf6232e187ea
----
77) 9b7f13b9b6d2b0d1623fd1e1aef67ecffeb363e9
units: 28100000
referrals:
----
78) 46b19c2e75e5d19772efc108d9d5f353047b2148
units: 281000000
referrals: ec931d7bef828d5beadac291490a71fae80a401b,
058f3f556ce5e6364dee07b85087de0896251544
----
79) 7412c8b906815c25c9031b8b3b9f4805015ac7f2
units: 28662000
referrals: 82cd4d648248322bf99d8cd587a63e77fd50b841
----
80) 69e1181e26fab42ab0f73797aa8e8b8dc77628cf
units: 240000000
referrals:
----
81) ec14389807f0b6499a032db73fd8a94c1204a216
units: 240000000
referrals:
----
82) 0759d0642cec4265b350c253889d1bf97539e10f
units: 296000000
referrals:
----
83) 25b78d4f814f12c6b2e1d87b28a56143e5322be4
units: 281000000
referrals: fe4ae074b84e9b13124da976be15d4f0b38685af
----
84) 40f341145b52905d0af17d3344fbe5ab8f47d679
units: 256412500
referrals: ed6f4cdc27eac292a2560b3eebf592598ad5ee67,
79d4d61060c4071cc8530b47ef7b93209a6d0c33
----
85) d2cb8b857b1aabe1f4d7404ebe260cca0862370d
units: 10680000
referrals:
----
86) 22f12615b4045a222393455b5aa9c1934071ed52
units: 250179000
referrals:
----
87) 17809259208a66d81e94b45c7014f791aacbbb0d
units: 251647500
referrals: 559bc75ec6e6ab2079421bbd0ebecad729a60fb2
----
88) 1eee9a2db92c46eaff06b3c31e34db7bd1ecddaf
units: 88920000
referrals:
----
89) 024e0b8e22743d89e0465d30c337f951918705fa
units: 296000000
referrals: 53cb6a0B87F0c438d033C59Ad757f6536564db28
----
90) 111484d3345848f43c226995174cda275e9d1d5c
units: 222610400
referrals:
----
91) 80847544a13209b76a70de0022be3dcaa193f88c
units: 281000000
referrals:
----
92) 61fc83ad0743fee0ca0b98aba94a7a9d0e50718a
units: 296000000
referrals: 6eeb9dffb820d13fc7ad5df4a0b97d4cebfa2a59,
6a7405cb055669af15d14392ed3ce4d330fd2cb5,
3e2ea46bdfe5ec7f1a12b3f24d47439e1b5e9ebf,
405481f2d0031115a14e3e3af0283cae429d93fb,
6c0b439b0f993202eda0a0496c4da8d23ad7c801,
024e0b8e22743d89e0465d30c337f951918705fa,
e8c762ae04480c230ee51f7caacf3c3c6a98a5ac,
1c0370b8711059cb73937b407fb18cf7bdb04f00,
78da7fca531d4b6710d616cdd63d77bc80e7fd32,
82094dfcd0e9d8c57c89a4ffd5491f7b16833fce,
4888ad1efed510667745edcc3e5b8fd4f40b2752,
7c4e6b55f58d5396b82d9128a705f14a36f4eb76,
e5756aa82a39d3e23dbc4831a8c6a6840b1663c2,
7946c0c8807b8458f32dcf6aad0ad6807907e6c5,
a6e26dbef099a1a4c99dd4c90159776c1dbf490b,
8d4b4aaf1f581c5528d006dc3458e76cf2c4555f,
fb50effc86d5bae52dd18578d9a0c26be5c0e8c2,
a3ec1ac881f6397d7ac0916c91013a159b2fa9a7,
b4748d7f8c30c637bf1c7a801e19be5cb2da7a8f,
e7fd8119ff58dc422da6c6ecfd0f8a16b4610062,
9f3d84a5f8f9beb9ef03257b511ebf02894a5955,
1d446853cb843c890895cc86b94204c44fdfc2b6,
67ae0ec54d93edec844c007bcae312102ee5f3f8,
2069fe5c5beaf44db81aed0a281512e40d3f809b,
850ec325a5cea0ab75acd402c9ae4dbf04b5ed9d,
b21d65b8dd5d419592915a1d9bf281b7036e4ca7,
2f5b28dd625ca194894165edaa4704ee75d4e4ab,
21fb0f62adbf72fb7b515552dfc4b243aef5439c,
5cbbd3c50dc65ffe4eeabae2b93eab52619306c7,
705039748ceb8f563110c18df60793265982ace2,
e8b9f63bb91ed1a2ab7d8a5073a5faf45c63340f,
b4a7f8acb4537a2a263f829b77fa58d1cd85b706,
312fc133481fdc7d08a1ecdb1ee3f0208cdebc39,
7c19a5ac994cf5bf2c1aed71d4def63d96ad88b8,
cb5d8b1c5f2564d541d0635576976c83842de95e,
30c45d0f4fba266731c0baa7b644d7854c10eeac,
30a23cfa9767f01dd562f8bcc46abfb3e45196a9,
8818518f5250ed3dc7bcf4e539f6b35b54e29b29,
abaeca5fd546ccfcaebb6fa1b69a6c6bdc9b44ec,
a4bb241c53eaa799071e6ad7b716dba89a441f69,
6eec71fc6195b0f835f19caa27b29461751128ce,
ed45a63726ebdb50573c8dee4e961fdbf20e0e4a,
ed116130ec54b9be4d58ac17a3849e3df497125a,
13da579450b27523006531dfcf8bb63500e4f871,
cf1d71b02aa523f72aca3f312505f01058772d45,
daab7eb7171aacd568127e246de64e28b4c6296e,
e0951e8a48456de9d31dae4b22c2e34af80d6800,
27370b2f0d172da1d506fa42efb4a71fc03f99af,
d463d867704d6e6356a09d9d8974c088cdd90809,
3ecbc92fa89b44de3b631553dbf62463e4534d75,
e2d0bafac71db8c6e4884c4dd010264ffa49c44b,
c1ed11d231e512982cd87f8b5861a0fed7a19153,
29cbccd01d8d2d3a4f4c437c4ae5039362df95e2,
bd793d8e9a131b9d31d8b3695676c0a408c0fdc4,
02107c1794C7074513874f4c608e2D00CdeE84f5,
E60110EdCDF0Db502b664F6F894f35B8C12414A9,
BE5be2be175E50D04F55BAF299E5f5B2062Ec934,
D4a3c4D6Cd564fdcAE0082f3A4419Bb7a4BfaE10,
c4ce1c95f0ed3ae7f1428712c1e768cef2462ec2,
f4fd96326fcF4E8b90C44fda94622F9c05F17DA9
----
93) 5aded2a2748e0dda11d5abbb0d423a94b828a5f0
units: 28100000
referrals:
----
94) 2aa4782c4cd8f26116b753016d3d789c45c077dc
units: 281000000
referrals:
----
95) 22719197dd5b49c4bce447230dc70f77e74020bc
units: 267000000
referrals:
----
96) 3603b37948c74b6e15c7988aae9624741e7b2398
units: 240300000
referrals:
----
97) 449b3b54c81883680e09f2197838c4d99e1d02ad
units: 24000000
referrals:
----
98) dc9d5636b830ce0b078dc2d546b9e1050ab60bc0
units: 28100000
referrals:
----
99) 1d3149339fdc6c70e0dd5de824d821848fdacd05
units: 267630000
referrals: b900548e21b6cbb11f578fca70dfb80cfbeb29d5,
b2f44b1a4f1ca5c7333151d8a730af6ac71ecd63,
546b76e990ad05f904c7e1ef648cab8234c7311c,
449b3b54c81883680e09f2197838c4d99e1d02ad,
ef98c81bd01bc0287d6abba8ad3d9496d93ded76,
23c8dc48eb57f7748d2ed89886581ff95978efd6
----
100) 9baadccbb7e843d3e4272cd74ee399c0c00f9954
units: 29600000
referrals:
----
101) ef98c81bd01bc0287d6abba8ad3d9496d93ded76
units: 240000000
referrals:
----
102) b863769e618825a4650bfb5ae8bf022bb92ad7e6
units: 87011400
referrals: 844950422e1f7306e536fa68a6eff106151cff24
----
103) dcf8023f18409005207026de7f0d1e780a94f742
units: 296000000
referrals: 5aded2a2748e0dda11d5abbb0d423a94b828a5f0,
5176fe6b79a153beba945f1ba0fb2fdf4f6a8dee,
b7d0fa0669730e46f9fae6d3a89bc26832b55dfe,
1eee9a2db92c46eaff06b3c31e34db7bd1ecddaf
----
104) 0850a184714490b7ef2204c83e69532080f4845f
units: 2960000
referrals:
----
105) 312fc133481fdc7d08a1ecdb1ee3f0208cdebc39
units: 56200000
referrals:
----
106) 982c8f4e601a63346b7bd0bede7c04e139e4f7bd
units: 281000000
referrals:
----
107) 8702ffdd9e3495a0bf81eb315f966284d41fb594
units: 125180400
referrals:
----
108) c531e47f882a34543cc77e92ffbba93d3ea11117
units: 198888300
referrals:
----
109) 36917bc886e619521d76f8994db962fc1e748e99
units: 296000000
referrals:
----
110) b2f44b1a4f1ca5c7333151d8a730af6ac71ecd63
units: 25365000
referrals:
----
111) 405481f2d0031115a14e3e3af0283cae429d93fb
units: 296000000
referrals: 72f72e91f658acfa46d6b6f2bf58ddd85193ec71
----
112) 82094dfcd0e9d8c57c89a4ffd5491f7b16833fce
units: 177600000
referrals: 16ce7d09d38b8e5dca13980cf9d4bce2d2244c77,
b4536dc4950b53e3c16fb85c5a56f9be239a7b6f,
deda2f75187f96965eb24fd53123f42a5ef342a8,
445bef61a9353db1813002375a305245dcbf367c,
b863769e618825a4650bfb5ae8bf022bb92ad7e6,
7245aa3377401c1a0482219053edc8545d65920f,
41e0d6fd4f722f39b3168e2cac4da2e0a2b8fe6b,
d9870a9677778db6da1da141d9faf3cc8c1dbe15,
0850a184714490b7ef2204c83e69532080f4845f,
ea9d1d1fa0b7bed82f0d2723883f0543a8a18e9a,
8b832f9201651e5f193f004997ccb507a19e2700,
65ebd58d35231c99c401da6b49deeb5e9be93748,
1d3149339fdc6c70e0dd5de824d821848fdacd05,
dc9d5636b830ce0b078dc2d546b9e1050ab60bc0,
8eecf761a6ac8076f2a1c23c4f87345b24e3f477,
941bbd3d449cd092781f6f12819dabd537c36947,
6390b88a6369d08e28ece4d741131d6ffbac224c,
dc89fc2f0d8b73d769c909eb04de30e235e7ad53
----
113) c6a3ee287edfce5bdef89121359fa63b50cb9cab
units: 156880000
referrals: 02714daebe7ec7c7803e6c2f5ac18814dae21bfe
----
114) 30c45d0f4fba266731c0baa7b644d7854c10eeac
units: 281000000
referrals:
----
115) c2dc0add26ee49fabf8e2908b9d5e0991317f922
units: 26700000
referrals:
----
116) 79edb79baecbc493157512fdc8771821548c637e
units: 296000
referrals: 61fc83ad0743fee0ca0b98aba94a7a9d0e50718a,
61fc83ad0743fee0ca0b98aba94a7a9d0e50718a
----
117) e0951e8a48456de9d31dae4b22c2e34af80d6800
units: 267000000
referrals:
----
118) 172aa24dcaefcb2ef14104761d37233ca6e6c166
units: 267000000
referrals:
----
119) 29cbccd01d8d2d3a4f4c437c4ae5039362df95e2
units: 267000000
referrals: 6782f058b8cc66625c9da50137cc71730bde1c08,
f1a51ea68acefc4395b5bc12a8c5d94f4d1a2470,
22719197dd5b49c4bce447230dc70f77e74020bc
----
120) 47a1147115756883038d000e526a4186902119e0
units: 136170000
referrals: d34a20e69403aa10913dd91455aee574f7611527
----
121) 3f947d3fd2ef078f77cf01f781a7392979dff02e
units: 203584500
referrals:
----
122) 014b5bb2f3800637e0e7de077a9af4803ab81e9c
units: 281000000
referrals:
----
123) cf1d71b02aa523f72aca3f312505f01058772d45
units: 281000000
referrals:
----
124) 3ecbc92fa89b44de3b631553dbf62463e4534d75
units: 267000000
referrals: b3cd225e39779c41dc994eb881b696f0cbc68e08,
40cc689036296a1d064805a5b913043a5177bc4e,
e443ee342ccd35b2ba7018aa9816abeaa8123586
----
125) 767868f8bf31973ca8f7559e5ea2cb17537003c5
units: 267000000
referrals:
----
126) 91c62ff6e40b9718af781148fac9d3b6684a4c24
units: 528960000
referrals:
----
127) 4ac0574bc2249b09808aa8f0dd4130b053889601
units: 2810000
referrals:
----
128) e2d0bafac71db8c6e4884c4dd010264ffa49c44b
units: 228819000
referrals:
----
129) 6789297f7623d5cd5a274df4e82dc0ddc0b96948
units: 80286900
referrals:
----
130) 65ebd58d35231c99c401da6b49deeb5e9be93748
units: 2960000
referrals:
----
131) a58f504430a6a2d969b959a73d590c54da79f800
units: 56200000
referrals:
----
132) 89728067216ca54361060316d766bf681853e7fa
units: 28100000
referrals:
----
133) 2e0f459670aa7119ed44da8b8358eb81b372ba3c
units: 27759990
referrals:
----
134) 58ed012303334209f945219fc00f707c0e2c5d80
units: 28100000
referrals: 5b5d95d32f55f08480704077888d713d50c8f18f
----
135) 91b91476849b46e0b6335c94302307f3e3348c85
units: 267000000
referrals:
----
136) 21a7619492f3d5aa58345556863026bba112db56
units: 117480000
referrals:
----
137) 365d913c0a29d709fe3e649a1376951661b96fd1
units: 267000000
referrals:
----
138) 0e6ce52c1d0257808be60c3f799172bd1ad30028
units: 267000000
referrals:
----
139) 632d80122ad73dc56cf7429dd038b514c25f32b3
units: 75027000
referrals:
----
140) c8219b660827326b3f55ea7776b6578a0a20d2cf
units: 102576000
referrals:
----
141) 44ee1512f5d7455523cda8338b04ae6a60a99de4
units: 58080000
referrals:
----
142) 789ab82821d9efb4511e3a7f75820037b560b768
units: 52652400
referrals:
----
143) 12ee88a73ab1d6609863f6c863fc26b017627542
units: 25852000
referrals:
----
144) 503c2c166b9a2785fa944d6d46b0247b1d1a5e8d
units: 296000000
referrals: 1e8c82755f62081cc7c760730677c5776da27c32
----
145) bb47df9c00cfc201e709764f31306d02220aa1b4
units: 5620000
referrals:
----
146) d281be420f04e2fa9f95172c94a8276c32f995fc
units: 267000000
referrals: a28e98b686096ac65474a27e9b3a80e4072e613f
----
147) f1a51ea68acefc4395b5bc12a8c5d94f4d1a2470
units: 267000000
referrals:
----
148) 7be8c96a98cb2fe2fd6c59edd2a83cebb3233ba2
units: 238965000
referrals:
----
149) 3a59dbb440dcb76b37c67c2891679270c5909e51
units: 48000000
referrals:
----
150) 16ce7d09d38b8e5dca13980cf9d4bce2d2244c77
units: 88800000
referrals:
----
151) 08fc58b846e69b332d6de7facba6062648490e8d
units: 182715000
referrals:
----
152) e67d86886f97daf5f059e876a4c3c04a514b3e69
units: 4496000
referrals:
----
153) f47ac69c3f6b7a2950523e641f09c80a0559d8e7
units: 281000000
referrals:
----
154) f81526ac03a5f38ff58eb9cd4b5caaa3d8ee6771
units: 252933000
referrals:
----
155) 68c950840858347b7f744f7c5926657e37fb4392
units: 267000000
referrals:
----
156) 423bdfeb5d25b0c8d5c6ff598e99eb883a340381
units: 52320000
referrals:
----
157) f6e03430b73670dc1bddc2520bebf08b3a79fe47
units: 24000000
referrals:
----
158) 6a7405cb055669af15d14392ed3ce4d330fd2cb5
units: 207200000
referrals: 994b046deb3d24b5b242d4ee07d00e1fed3f3261,
c6a3ee287edfce5bdef89121359fa63b50cb9cab,
36917bc886e619521d76f8994db962fc1e748e99,
503c2c166b9a2785fa944d6d46b0247b1d1a5e8d,
717d218fbc3d0fe0997597d772848d328a4bb2a0,
eb87c91742da3adb33c8f0faf0c0111f1bdb6b1a,
adc0b7a58ba5fbd9e8a6980350b39c52704a2f6a,
ef4af552b493cd52ca79765ac29098e5b65c3989,
c691c060ebeeab5655cc997246a962e7162e1e8e,
8702ffdd9e3495a0bf81eb315f966284d41fb594,
41224e44922d2d98aa6360c705772c59aacec96b
----
159) 8ae39c6ab7f21e06b21e8d725f87f604be371ff9
units: 296000000
referrals: 8c4e1d2cedaa4c017269592fc70aaf20a9b48ada,
2ea252ba869eed2a248240a00207583f934404ad,
8cb5d5bf581af4f3924adbf69a08db762b946ce2,
4d95ef6e08fb61207ef4d45b79d8b150d9382c7e,
2387dbc688e1f53b3d01a23a4d7ec4577c7f393b
----
160) 41e0d6fd4f722f39b3168e2cac4da2e0a2b8fe6b
units: 29600000
referrals:
----
161) 8ad7520f95f76af7328f8a45baa5d26ac1c79098
units: 281000000
referrals: 07bc70d8b22c9431ed35b9007c7646c77053ca75
----
162) 9650a79fc1bccbd27dfc0dfb76769d15c54472b9
units: 203454000
referrals:
----
163) 6c0b439b0f993202eda0a0496c4da8d23ad7c801
units: 296000000
referrals: 9baadccbb7e843d3e4272cd74ee399c0c00f9954,
a23344184e2c6e1aa744e9b7cd4b513cea5afb3c,
70d79418a686ee65a449c2d06220e643155f0f14,
ee5cae74e334c2f1af65d19052b202c2f6bd4ab4,
f2869c98fb3d1753346449b06c18f27dc6611292,
982c8f4e601a63346b7bd0bede7c04e139e4f7bd,
26bc99a54c290a7865dcbf0b1ed8a4f013ea964f,
89728067216ca54361060316d766bf681853e7fa,
111484d3345848f43c226995174cda275e9d1d5c,
7e0d4c8aa8ba9e422cd4414c5e0a9e720d4d2d46,
767868f8bf31973ca8f7559e5ea2cb17537003c5,
68c950840858347b7f744f7c5926657e37fb4392,
5159f99d35cb74796de5664da456b79a21ac5449
----
164) 30a23cfa9767f01dd562f8bcc46abfb3e45196a9
units: 168600000
referrals:
----
165) a111006ccb092aa6a4044bc322135d8b0842a9f2
units: 29645500
referrals:
----
166) 21fb0f62adbf72fb7b515552dfc4b243aef5439c
units: 296000000
referrals:
----
167) 4888ad1efed510667745edcc3e5b8fd4f40b2752
units: 150960000
referrals:
----
168) 880a9d891b285650defe2969ed22608d26ddfff3
units: 40050000
referrals:
----
169) 04559347c9c78805cb65f6b290c0f6fb9fd2c3fd
units: 592000000
referrals:
----
170) f92e87d99190a1bab182c2f1014738f58f58a3e5
units: 31472000
referrals:
----
171) b4a7f8acb4537a2a263f829b77fa58d1cd85b706
units: 281000000
referrals: c979748097d71aa9cf71ad229608c5de0a0965a4
----
172) f6389ce8ff6b83c0a099206461317bc2bca10ef5
units: 296000000
referrals:
----
173) cb5d8b1c5f2564d541d0635576976c83842de95e
units: 140500000
referrals:
----
174) c1ed11d231e512982cd87f8b5861a0fed7a19153
units: 267000000
referrals: d281be420f04e2fa9f95172c94a8276c32f995fc,
365d913c0a29d709fe3e649a1376951661b96fd1
----
175) a6e9ce90646c5624c7cd530422eaa3a6662f1992
units: 24000000
referrals:
----
176) fa01965486e0b6bfe4f314b3c1d7914abc0d2f80
units: 142672000
referrals:
----
177) 7c19a5ac994cf5bf2c1aed71d4def63d96ad88b8
units: 28100000
referrals:
----
178) 8c5a840b5a9c2d78f665c93c63c3061f187c3346
units: 140500000
referrals:
----
179) 82cd4d648248322bf99d8cd587a63e77fd50b841
units: 28071900
referrals:
----
180) b6165e06be0cf63bd008c98000f933e365b09f08
units: 51142000
referrals:
----
181) d463d867704d6e6356a09d9d8974c088cdd90809
units: 267000000
referrals:
----
182) 8d4b4aaf1f581c5528d006dc3458e76cf2c4555f
units: 296000000
referrals: 1c2Cc5C54F3362Ff2bFDc03c0B1f9664b2073563
----
183) 67ae0ec54d93edec844c007bcae312102ee5f3f8
units: 249232000
referrals:
----
184) 705039748ceb8f563110c18df60793265982ace2
units: 296000000
referrals: f6389ce8ff6b83c0a099206461317bc2bca10ef5,
4ef240f6c76f73b1ac2d068f0309987a9f98a9e3
----
185) dc9e7a6a086dbb784c7bd7d84b4884b0ea1f168a
units: 14050000
referrals: 093ff2fb6b3c815ff955248519ec7b7af42e665e
----
186) daab7eb7171aacd568127e246de64e28b4c6296e
units: 281000000
referrals: 3f947d3fd2ef078f77cf01f781a7392979dff02e
----
187) e5756aa82a39d3e23dbc4831a8c6a6840b1663c2
units: 14800000
referrals:
----
188) e6efe00342d70cb85be60ab796ab98cc2a71f23b
units: 275099000
referrals:
----
189) 5f25a86383888f74b944576c92d4ec75b35ea66d
units: 16682160
referrals:
----
190) 4d95ef6e08fb61207ef4d45b79d8b150d9382c7e
units: 61410000
referrals:
----
191) 49ec0ed807417b69e2950fa809a60e1994d2dc6b
units: 44880000
referrals:
----
192) 07bc70d8b22c9431ed35b9007c7646c77053ca75
units: 24000000
referrals:
----
193) b900548e21b6cbb11f578fca70dfb80cfbeb29d5
units: 36060000
referrals:
----
194) 13da579450b27523006531dfcf8bb63500e4f871
units: 281000000
referrals:
----
195) aaf4394afbe1b7c6c8379da08a999da788a9728c
units: 26700000
referrals: 0441fd8f6b6a9fee71015f5e90a564bf79d39052
----
196) e7e25b26d7d16689a4ba3f5410b85e30bd297d39
units: 58785200
referrals: 2e0f459670aa7119ed44da8b8358eb81b372ba3c
----
197) c691c060ebeeab5655cc997246a962e7162e1e8e
units: 56200000
referrals: 128C688929B28fFD468705220EcDd71A74E81B89
----
198) f4245a2ca2faa8f21dbda9aaf8213b45a8ff008c
units: 281000000
referrals:
----
199) 8cb5d5bf581af4f3924adbf69a08db762b946ce2
units: 267000000
referrals:
----
200) 871d583deb3083993b28487beb934eef155e0118
units: 96387000
referrals:
----
201) 612772ec0968c7bf79856ca3288ef383ec90da9d
units: 239766000
referrals:
----
202) 7946c0c8807b8458f32dcf6aad0ad6807907e6c5
units: 292744000
referrals:
----
203) b4748d7f8c30c637bf1c7a801e19be5cb2da7a8f
units: 296000000
referrals: ee48960e29c896a4726a17ec65a9dd39555476b0,
abe2d1bbc5833da77af5e79ee08dba4b9acd2114,
40f341145b52905d0af17d3344fbe5ab8f47d679,
178132f6c553ce96d4b47ec224a5c37e76be39dd,
c8219b660827326b3f55ea7776b6578a0a20d2cf
----
204) 717d218fbc3d0fe0997597d772848d328a4bb2a0
units: 296000000
referrals: f4245a2ca2faa8f21dbda9aaf8213b45a8ff008c,
93967805faa6eabe9de62d578b1f076863908133
----
205) 5cbbd3c50dc65ffe4eeabae2b93eab52619306c7
units: 296000000
referrals: 46b19c2e75e5d19772efc108d9d5f353047b2148
----
206) ed45a63726ebdb50573c8dee4e961fdbf20e0e4a
units: 281000000
referrals:
----
207) e53b3cb82815b613062dc657a9fcd6cb77ec595a
units: 267000000
referrals:
----
208) 08eb3f0457bf0de318f46b886df6805fc788c1d2
units: 53133000
referrals:
----
209) a28e98b686096ac65474a27e9b3a80e4072e613f
units: 80100000
referrals:
----
210) 21462509020a90abc8864c02dcb43348bcce63ec
units: 296000
referrals:
----
211) 475131c7992a6666969b344ad8cb8b9f8e98e606
units: 143640000
referrals:
----
212) 3023138cb3aeeae67c4033015a39b56af582e3fc
units: 265665000
referrals: 17809259208a66d81e94b45c7014f791aacbbb0d
----
213) a576639a4884fb97648111d0cafe9bc89ff5436a
units: 113960000
referrals: fa01965486e0b6bfe4f314b3c1d7914abc0d2f80
----
214) fcbdb50e1d99310380bfab314a66a0f4e2b23875
units: 14800000
referrals: 8dd3382eb29c4eccf48482d7d7ed93301638b702
----
215) cec4df3865be446e9e8d2fd9cb9e182f698c6e1b
units: 267000000
referrals:
----
216) 8e367218246f9df464dd925169c470d8f0622222
units: 158865000
referrals:
----
217) ee5cae74e334c2f1af65d19052b202c2f6bd4ab4
units: 296000000
referrals:
----
218) 2069fe5c5beaf44db81aed0a281512e40d3f809b
units: 29600000
referrals:
----
219) 93a67e0b06f9dc0fdab8a041c34b74fb29c018f8
units: 59010000
referrals:
----
220) 5c386ad5f389242c066f19bc3b423a6097e9efb7
units: 140500000
referrals:
----
221) ed6f4cdc27eac292a2560b3eebf592598ad5ee67
units: 235224000
referrals:
----
222) a6e26dbef099a1a4c99dd4c90159776c1dbf490b
units: 296000000
referrals: ed699e481adb3111f23742f0646a5c7695eebd10
----
223) fe4ae074b84e9b13124da976be15d4f0b38685af
units: 281000000
referrals:
----
224) fc6d028371e3eced057f7b8a8a57f7c68ad306d4
units: 112140000
referrals:
----
225) 2387dbc688e1f53b3d01a23a4d7ec4577c7f393b
units: 161802000
referrals:
----
226) 61229a2ac51b2430d228f58a2658e1770346f0f6
units: 247242000
referrals: 22f12615b4045a222393455b5aa9c1934071ed52,
503bb794fe7a57a15fb7e7967af22bf341682c5b
----
227) e8b9f63bb91ed1a2ab7d8a5073a5faf45c63340f
units: 19670000
referrals:
----
228) 12cbbf890b773fc9c4990aa3c3333267c14dd9f8
units: 219750000
referrals: 14e989718c622e6a90531c515bf2c9ac5fb8d713,
5dce7412dac1860b1c9eaceae3a09ed2f4ec7600,
6db0ba283092dbb17e30639193dbf67091d16ad0,
9944aC84145836C679acA9968992E12A9330Ed2c,
e67d86886f97daf5f059e876a4c3c04a514b3e69,
0495e62e568d477c10de7116c8938350e2da690e,
53daccec0fda34c30f89cd64ad29a9bbe84d275e,
58ed012303334209f945219fc00f707c0e2c5d80,
b07a67678bab935c2ce60edc41463aa5785e9dd4
----
229) 14e989718c622e6a90531c515bf2c9ac5fb8d713
units: 188650000
referrals:
----
230) a4bb241c53eaa799071e6ad7b716dba89a441f69
units: 281000000
referrals: 8c5a840b5a9c2d78f665c93c63c3061f187c3346
----
231) 0495e62e568d477c10de7116c8938350e2da690e
units: 56200000
referrals:
----
232) 6e498f21f09919914f5cc486f7be690e2bebd9e7
units: 281000000
referrals:
----
233) 0730509fac1c163e947c26cfa8a40fe76cc846d4
units: 59200000
referrals: 072a9b89f4462d3c2f9183e82027b42299b74fd9,
d0a26322abfaa561da11713e76af4e5611b9a4fd,
352edba41626f7a6c6641d71ba94a5e28d49b7a9,
91e0c71217225a4870a6ef35a7ca0b674941a746,
899530f3464a37920674b536384c299c6a09eeee,
48e36BdBBc54219136FDd20685ccC73166Fe8324,
ca9e4ec800b85345b186dcda5c02671a861dc067,
234cc5b3c115e25efc9e3b9c5bc8bfc6b69cb5e6,
Fc6D028371e3eCED057f7B8A8A57f7C68ad306D4,
21A7619492f3D5aa58345556863026bBa112Db56
----
234) 8d24e10d251a575cebe15a18c8f2e9e956f57b60
units: 58904000
referrals:
----
235) 25e7d4c6529dc774ae1a44e686084273dab679fb
units: 272289000
referrals: 5db0ce4be877074f838ec85eb196784ee529df15,
cec4df3865be446e9e8d2fd9cb9e182f698c6e1b,
91b91476849b46e0b6335c94302307f3e3348c85
----
236) bd793d8e9a131b9d31d8b3695676c0a408c0fdc4
units: 133500000
referrals:
----
237) 994b046deb3d24b5b242d4ee07d00e1fed3f3261
units: 296000000
referrals: a547fdd9210fac0786d05af03fb289d1947531d4,
a576639a4884fb97648111d0cafe9bc89ff5436a,
9650a79fc1bccbd27dfc0dfb76769d15c54472b9
----
238) 03fd9c6a9fd55eea203d2f1e5e6e33ecc42298a6
units: 80901000
referrals:
----
239) eb87c91742da3adb33c8f0faf0c0111f1bdb6b1a
units: 281000000
referrals: 2aa4782c4cd8f26116b753016d3d789c45c077dc
----
240) 1d446853cb843c890895cc86b94204c44fdfc2b6
units: 29600000
referrals:
----
241) f2869c98fb3d1753346449b06c18f27dc6611292
units: 37888000
referrals:
----
242) 78a607231617d6e0e3719ed3e127e5084f6cb1b6
units: 46560000
referrals:
----
243) 70d79418a686ee65a449c2d06220e643155f0f14
units: 296000000
referrals:
----
244) 290dbb38d36e12dffa7ea95d1dd983caf842db61
units: 9176000
referrals:
----
245) abe2d1bbc5833da77af5e79ee08dba4b9acd2114
units: 59200000
referrals:
----
246) 02714daebe7ec7c7803e6c2f5ac18814dae21bfe
units: 281000000
referrals: 25b78d4f814f12c6b2e1d87b28a56143e5322be4,
f47ac69c3f6b7a2950523e641f09c80a0559d8e7,
4877f396950c4c93dcf29ab78f5a246548aaf40e,
c2dc0add26ee49fabf8e2908b9d5e0991317f922,
6890eef74ea683a9383b6a32099b3967eae7907e,
0e6ce52c1d0257808be60c3f799172bd1ad30028,
612772ec0968c7bf79856ca3288ef383ec90da9d
----
247) ed699e481adb3111f23742f0646a5c7695eebd10
units: 267000000
referrals:
----
248) c3646c9f183a7ff16addb0b778221519a1faeac4
units: 134760000
referrals:
----
249) a23344184e2c6e1aa744e9b7cd4b513cea5afb3c
units: 29600000
referrals:
----
250) abaeca5fd546ccfcaebb6fa1b69a6c6bdc9b44ec
units: 31472000
referrals: 7412c8b906815c25c9031b8b3b9f4805015ac7f2,
b6165e06be0cf63bd008c98000f933e365b09f08
----
251) eec9aed9d7d5430dbee4f6425c5974a6dce5dd78
units: 24000000
referrals:
----
252) c555acf73a1164f7f2eee5e7f59fee502f17afec
units: 240000000
referrals: ec14389807f0b6499a032db73fd8a94c1204a216
----
253) 7c4e6b55f58d5396b82d9128a705f14a36f4eb76
units: 242720000
referrals:
----
254) ee48960e29c896a4726a17ec65a9dd39555476b0
units: 275840000
referrals:
----
255) 850ec325a5cea0ab75acd402c9ae4dbf04b5ed9d
units: 148000000
referrals:
----
256) 81b01fd31017db4a6d85c098c26902ba15e29625
units: 281000000
referrals:
----
257) 6eec71fc6195b0f835f19caa27b29461751128ce
units: 28100000
referrals:
----
258) dc89fc2f0d8b73d769c909eb04de30e235e7ad53
units: 3296297
referrals:
----
259) 6eeb9dffb820d13fc7ad5df4a0b97d4cebfa2a59
units: 296000000
referrals:
----
260) 85dcfd9d342232e008c13ea3ef47d3532115a92a
units: 5624000
referrals:
----
261) 352edba41626f7a6c6641d71ba94a5e28d49b7a9
units: 73425000
referrals: 789ab82821d9efb4511e3a7f75820037b560b768
----
262) c64c8e600a35d3786cf0b13a46744e010617b080
units: 2696000
referrals: 7863b0abf3926301b380c01bf7e0ff08133be6cc
----
263) 1c0370b8711059cb73937b407fb18cf7bdb04f00
units: 85248000
referrals: 8ae39c6ab7f21e06b21e8d725f87f604be371ff9,
c97c55442099a325ed7b399925846b51d60ead3b,
a58f504430a6a2d969b959a73d590c54da79f800,
93a67e0b06f9dc0fdab8a041c34b74fb29c018f8,
a111006ccb092aa6a4044bc322135d8b0842a9f2,
e75727bd26e24347d01c33b134d77d0d9bcc6b87
----
264) e75727bd26e24347d01c33b134d77d0d9bcc6b87
units: 191973000
referrals:
----
265) 0441fd8f6b6a9fee71015f5e90a564bf79d39052
units: 6141000
referrals:
----
266) 7863b0abf3926301b380c01bf7e0ff08133be6cc
units: 296000
referrals:
----
267) 27370b2f0d172da1d506fa42efb4a71fc03f99af
units: 52332000
referrals: 47a1147115756883038d000e526a4186902119e0,
632d80122ad73dc56cf7429dd038b514c25f32b3,
f6e03430b73670dc1bddc2520bebf08b3a79fe47
----
268) 91e0c71217225a4870a6ef35a7ca0b674941a746
units: 101460000
referrals:
----
269) 3e091d480d0e3429fab8e072075b7502f218ebb8
units: 260592000
referrals:
----
270) 072a9b89f4462d3c2f9183e82027b42299b74fd9
units: 59200000
referrals: 9b7f13b9b6d2b0d1623fd1e1aef67ecffeb363e9,
61178ed786321f6acdf01633e336d94fc42ec53a,
5f25a86383888f74b944576c92d4ec75b35ea66d,
d2cb8b857b1aabe1f4d7404ebe260cca0862370d,
08eb3f0457bf0de318f46b886df6805fc788c1d2,
621cb7f95cd54b6884ca331cfb5e610fa5cb092d,
3a59dbb440dcb76b37c67c2891679270c5909e51
----
271) fb50effc86d5bae52dd18578d9a0c26be5c0e8c2
units: 296000000
referrals: 6e498f21f09919914f5cc486f7be690e2bebd9e7
----
272) c97c55442099a325ed7b399925846b51d60ead3b
units: 59572000
referrals: 43234e1a1c49023725af693b6bc5a76c11c46765
----
273) c979748097d71aa9cf71ad229608c5de0a0965a4
units: 55919000
referrals: e7e25b26d7d16689a4ba3f5410b85e30bd297d39,
1689ab3994e413046f3d8e153768c7757a015420
----
274) 559bc75ec6e6ab2079421bbd0ebecad729a60fb2
units: 231840000
referrals:
----
275) a547fdd9210fac0786d05af03fb289d1947531d4
units: 296000000
referrals:
----
276) e8c762ae04480c230ee51f7caacf3c3c6a98a5ac
units: 296000000
referrals: 0759d0642cec4265b350c253889d1bf97539e10f,
e8a1a820cdc94ca5e2726612a8dc0e7b0e33adaa,
80847544a13209b76a70de0022be3dcaa193f88c,
81b01fd31017db4a6d85c098c26902ba15e29625,
e6efe00342d70cb85be60ab796ab98cc2a71f23b
----
277) fad9ca41589b3657d56e1a6b38601dcbc5dab508
units: 88504000
referrals:
----
278) 8c4e1d2cedaa4c017269592fc70aaf20a9b48ada
units: 190920000
referrals:
----
279) 8eecf761a6ac8076f2a1c23c4f87345b24e3f477
units: 30067000
referrals:
----
280) 4877f396950c4c93dcf29ab78f5a246548aaf40e
units: 267000000
referrals:
----
281) 7e0d4c8aa8ba9e422cd4414c5e0a9e720d4d2d46
units: 267000000
referrals:
----
282) 02107c1794c7074513874f4c608e2d00cdee84f5
units: 4440000
referrals: fcbdb50e1d99310380bfab314a66a0f4e2b23875,
89cbf332571f54e2d1ceaddc2b91b6a96f21df00
----
283) 8b832f9201651e5f193f004997ccb507a19e2700
units: 296000000
referrals: 12cbbf890b773fc9c4990aa3c3333267c14dd9f8
----
284) 72f72e91f658acfa46d6b6f2bf58ddd85193ec71
units: 41440000
referrals:
----
285) 234cc5b3c115e25efc9e3b9c5bc8bfc6b69cb5e6
units: 267000000
referrals: ebacba195accc94e6ac9b34207f3916c41119329
----
286) 445bef61a9353db1813002375a305245dcbf367c
units: 20720000
referrals: 96bad93f26c64e72f7b22e3dee731ca51e84c78b,
12ee88a73ab1d6609863f6c863fc26b017627542
----
287) E60110EdCDF0Db502b664F6F894f35B8C12414A9
units: 12100000
referrals:
----
288) BE5be2be175E50D04F55BAF299E5f5B2062Ec934
units: 12100000
referrals:
----
289) D4a3c4D6Cd564fdcAE0082f3A4419Bb7a4BfaE10
units: 24200000
referrals:
----
290) f4fd96326fcF4E8b90C44fda94622F9c05F17DA9
units: 4840000
referrals:
----
291) 53cb6a0B87F0c438d033C59Ad757f6536564db28
units: 24200000
referrals:
----
`

	founders := []string{
		strings.ToLower("8A85c533693a87837380d9225d226e334663d104"),
		strings.ToLower("EF626c6425A2b077c23c2d747FCfE65777F66B10"),
		strings.ToLower("ceaE30276B9fD5FA44366167e64728180eb3962c"),
		strings.ToLower("1349DCDd92BA65Cf2234eD8c61C72DdF1f95400E"),
		strings.ToLower("acd745EB1F708C323C2167966fcA4503430705E1"),
		strings.ToLower("13B7fD960C3c105c0a80f05a2430783345A7c8dC"),
	}

	levelToPercent := map[uint8]uint64{
		0: 800,
		1: 400,
		2: 150,
		3: 50,
		4: 50,
		5: 25,
		6: 25,
	}

	// parse the data:
	dataMap := parseData(data)

	// make the units uint64:
	updated := map[string]treeUint{}
	for address, oneTree := range dataMap {
		casted, err := strconv.Atoi(oneTree.units)
		if err != nil {
			panic(err)
		}

		updated[address] = treeUint{
			units:     uint64(casted),
			referrals: oneTree.referrals,
			top:       0,
			referred:  0,
		}
	}

	// initialize the founders if not in the parsed list:
	for _, oneAddress := range founders {
		if _, ok := updated[oneAddress]; !ok {
			updated[oneAddress] = treeUint{
				units:     0,
				referrals: []string{},
				top:       0,
				referred:  0,
			}
		}
	}

	// calculate the referral amount and build the final tree:
	for _, oneTree := range updated {
		// for each referral:
		for _, oneReferralAddress := range oneTree.referrals {
			// execute the referral: calculation:
			referredCalculation, founderCalculation := calculateTreeValue(
				updated[oneReferralAddress].units,
				oneReferralAddress,
				updated,
				levelToPercent,
				founders,
				0,
				map[string]uint64{},
				map[string]uint64{},
			)

			// referrals:
			for oneAddress, oneValue := range referredCalculation {
				updated[oneAddress] = treeUint{
					units:     updated[oneAddress].units,
					referred:  updated[oneAddress].referred + oneValue,
					top:       updated[oneAddress].top,
					referrals: updated[oneAddress].referrals,
				}
			}

			// founders:
			for oneAddress, oneValue := range founderCalculation {
				updated[oneAddress] = treeUint{
					units:     updated[oneAddress].units,
					referred:  updated[oneAddress].referred,
					top:       updated[oneAddress].top + oneValue,
					referrals: updated[oneAddress].referrals,
				}
			}
		}
	}

	// compound the total units without referrals or top:
	totalWithoutReferrals := uint64(0)
	for _, oneTree := range updated {
		totalWithoutReferrals += oneTree.units
	}

	// compound the total:
	total := uint64(0)
	for _, oneTree := range updated {
		total += oneTree.units + oneTree.top + oneTree.referred
	}

	// calculate the total with cents:
	totalWithCentsTree := map[string]treeUint{}
	ratio := float64(totalWithoutReferrals) / float64(total)
	multiplier := uint64(float64(100000) * ratio)
	for address, oneTree := range updated {
		totalWithCentsTree[address] = treeUint{
			units:     oneTree.units * multiplier,
			referred:  oneTree.referred * multiplier,
			top:       oneTree.top * multiplier,
			referrals: oneTree.referrals,
		}
	}

	totalWithCents := uint64(0)
	for _, oneTree := range totalWithCentsTree {
		totalWithCents += oneTree.units
	}

	// compound the total referred:
	totalReferredWithCents := uint64(0)
	for _, oneTree := range totalWithCentsTree {
		totalReferredWithCents += oneTree.referred
	}

	// compound the total referred:
	totalFoundersCents := uint64(0)
	for _, oneTree := range totalWithCentsTree {
		totalFoundersCents += oneTree.top
	}

	// add the 15M to the founders:
	partPerFounder := uint64(totalEmissionWithCents * founderShare / len(founders))
	for _, oneAddress := range founders {
		totalWithCentsTree[oneAddress] = treeUint{
			units:     totalWithCentsTree[oneAddress].units,
			referred:  totalWithCentsTree[oneAddress].referred,
			top:       totalWithCentsTree[oneAddress].top + partPerFounder,
			referrals: totalWithCentsTree[oneAddress].referrals,
		}
	}

	// compound the total for founders:
	totalFounderWithCents := uint64(0)
	for _, oneTree := range totalWithCentsTree {
		totalFounderWithCents += oneTree.top
	}

	totalBackToDAOWithCents := totalEmissionWithCents - (totalWithCents + totalFounderWithCents + totalReferredWithCents)

	fmt.Printf("\ntotal include cents, the token is divisible in 1/100.000.000\n------\n")
	fmt.Printf("\ntotal to contributors: %d\n", totalWithCents)
	fmt.Printf("\ntotal back to DAO: %d\n", totalBackToDAOWithCents)
	fmt.Printf("\ntotal founders: %d \n", totalFounderWithCents)
	fmt.Printf("\ntotal referred: %d \n", totalFoundersCents)
	fmt.Printf("\ntotal emission: %d \n", totalBackToDAOWithCents+totalWithCents+totalReferredWithCents+totalFounderWithCents)
	fmt.Printf("\nreferring percentage: %f \n\n", float64(totalFoundersCents+totalReferredWithCents)/float64(totalWithCents))

	for oneAddress, oneTree := range totalWithCentsTree {
		referralListStr := "-"
		if len(oneTree.referrals) > 0 {
			referralListStr = "<ol>"
			for _, oneReferral := range oneTree.referrals {
				referralListStr = fmt.Sprintf("%s<li>%s</li>", referralListStr, oneReferral)
			}

			referralListStr = fmt.Sprintf("%s%s", referralListStr, "</ol>")

		}

		fmt.Printf("| %s | %s | %d | %d | %d |\n", oneAddress, referralListStr, oneTree.units, oneTree.referred, oneTree.top)
	}

}

func calculateTreeValue(
	amount uint64,
	owner string,
	data map[string]treeUint,
	levelToPercent map[uint8]uint64,
	founders []string,
	level uint8,
	referredCalculation map[string]uint64,
	founderCalculation map[string]uint64,
) (map[string]uint64, map[string]uint64) {
	if _, ok := levelToPercent[level]; !ok {
		return referredCalculation, founderCalculation
	}

	// create the output vars:
	outputReferred := referredCalculation
	outputFounder := founderCalculation

	// find the parent:
	parentAddress := findParentAddressIfAny(owner, data)
	if parentAddress != "" {
		// add the direct referred value to the owner:
		amountToAdd := (amount * levelToPercent[level]) / clevelDivider
		outputReferred[parentAddress] += amountToAdd

		// continue next level:
		return calculateTreeValue(
			amount,
			parentAddress,
			data,
			levelToPercent,
			founders,
			level+1,
			outputReferred,
			outputFounder,
		)
	}

	// no parent, so split the remaining referral to founders:
	additionalFounder := splitToFounder(amount, outputReferred, levelToPercent, founders)
	for oneAddress, oneValue := range additionalFounder {
		outputFounder[oneAddress] += oneValue
	}

	return outputReferred, outputFounder
}

func splitToFounder(
	amountUints uint64,
	referred map[string]uint64,
	levelToPercent map[uint8]uint64,
	founders []string,
) map[string]uint64 {
	output := map[string]uint64{}
	amountInTree := len(referred)
	remainingLevels := len(levelToPercent) - amountInTree
	if remainingLevels <= 0 {
		return output
	}

	for i := 0; i < remainingLevels; i++ {
		newLevel := uint8(amountInTree + i)
		value := (amountUints * levelToPercent[newLevel]) / clevelDivider
		partForEach := value / uint64(len(founders))
		for _, oneFounderAddress := range founders {
			if _, ok := output[oneFounderAddress]; ok {
				output[oneFounderAddress] += partForEach
				continue
			}

			output[oneFounderAddress] = partForEach
		}

	}

	return output
}

func findParentAddressIfAny(address string, data map[string]treeUint) string {
	for oneAddress, oneTree := range data {
		for _, oneReferral := range oneTree.referrals {
			if oneReferral == address {
				return oneAddress
			}
		}
	}

	return ""
}

// parseData fills a map of tree structs based on the provided data
func parseData(data string) map[string]tree {
	result := map[string]tree{}
	blocks := strings.Split(data, "----")

	for _, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		if len(lines) <= 0 {
			continue
		}

		var address, units string
		referrals := []string{}
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if strings.Contains(line, ")") {
				index := strings.Index(line, ")")
				address = strings.TrimSpace(line[index+1:])
				continue
			}

			if strings.HasPrefix(line, "units:") {
				units = strings.TrimSpace(strings.TrimPrefix(line, "units:"))
				continue
			}

			if strings.HasPrefix(line, "referrals:") {
				referrals = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "referrals:")), ", ")
				continue
			}

			if strings.HasSuffix(line, ",") {
				index := strings.Index(line, ",")
				referrals = append(referrals, line[:index])
				continue
			}

			referrals = append(referrals, line)
		}

		address = strings.TrimSuffix(address, ",")
		if address == "" || units == "" {
			continue
		}

		cleaned := []string{}
		for _, oneReferral := range referrals {
			trimmed := strings.TrimSpace(oneReferral)
			trimmed = strings.TrimSuffix(trimmed, ",")
			if trimmed != "" {
				cleaned = append(cleaned, strings.ToLower(trimmed))
			}
		}

		result[strings.ToLower(address)] = tree{
			units:     units,
			referrals: cleaned,
		}
	}

	return result
}
