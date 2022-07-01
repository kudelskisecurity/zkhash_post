
package griffin

import (
	"math/big"
	"github.com/consensys/gnark/frontend"
)

// Circuit defines the Griffin-Pi permutation
// https://eprint.iacr.org/2022/403
// with parameters t = 3, d = 5, R = 12 for the
// curve BLS-12-381

type Circuit struct {
	X0 frontend.Variable `gnark:"x"`
	X1 frontend.Variable `gnark:"x"`
	X2 frontend.Variable `gnark:"x"`

	Y0 frontend.Variable `gnark:",public"`
	Y1 frontend.Variable `gnark:",public"`
	Y2 frontend.Variable `gnark:",public"`
}

func exp_d(api frontend.API, x frontend.Variable) frontend.Variable {

	// Step 1: t5 = x^0x2
	t5 := api.Mul(x, x)

	// Step 2: t7 = x^0x4
	t7 := api.Mul(t5, t5)

	// Step 3: t6 = x^0x5
	t6 := api.Mul(x, t7)

	return t6

}

/* Generated via https://github.com/mmcloughlin/addchain
   addition chain for d_inv = 5^{-1} mod p -1 */

func add_chain_exp_d_inv(api frontend.API, x frontend.Variable) frontend.Variable {

	// Step 1: t5 = x^0x2
	t5 := api.Mul(x, x)

	// Step 2: t7 = x^0x4
	t7 := api.Mul(t5, t5)

	// Step 3: t6 = x^0x5
	t6 := api.Mul(x, t7)

	// Step 4: t9 = x^0x7
	t9 := api.Mul(t5, t6)

	// Step 5: t0 = x^0xa
	t0 := api.Mul(t6, t6)

	// Step 6: t1 = x^0xb
	t1 := api.Mul(x, t0)

	// Step 7: z = x^0xd
	z := api.Mul(t5, t1)

	// Step 8: t3 = x^0xf
	t3 := api.Mul(t5, z)

	// Step 9: t8 = x^0x11
	t8 := api.Mul(t5, t3)

	// Step 10: t2 = x^0x13
	t2 := api.Mul(t5, t8)

	// Step 11: t14 = x^0x17
	t14 := api.Mul(t7, t2)

	// Step 12: t12 = x^0x21
	t12 := api.Mul(t0, t14)

	// Step 13: t13 = x^0x23
	t13 := api.Mul(t5, t12)

	// Step 14: t11 = x^0x25
	t11 := api.Mul(t5, t13)

	// Step 15: t4 = x^0x27
	t4 := api.Mul(t5, t11)

	// Step 16: t10 = x^0x2b
	t10 := api.Mul(t7, t4)

	// Step 17: t16 = x^0x2f
	t16 := api.Mul(t7, t10)

	// Step 18: t0 = x^0x33
	t0 = api.Mul(t7, t16)

	// Step 19: t15 = x^0x35
	t15 := api.Mul(t5, t0)

	// Step 20: t5 = x^0x39
	t5 = api.Mul(t7, t15)

	// Step 21: t7 = x^0x3d
	t7 = api.Mul(t7, t5)

	// Step 22: t17 = x^0x5c
	t17 := api.Mul(t13, t5)

	// Step 28: t17 = x^0x1700
	for s := 0; s < 6; s++ {
		t17 = api.Mul(t17, t17)
	}

	// Step 29: t16 = x^0x172f
	t16 = api.Mul(t16, t17)

	// Step 35: t16 = x^0x5cbc0
	for s := 0; s < 6; s++ {
		t16 = api.Mul(t16, t16)
	}

	// Step 36: t16 = x^0x5cbe1
	t16 = api.Mul(t12, t16)

	// Step 42: t16 = x^0x172f840
	for s := 0; s < 6; s++ {
		t16 = api.Mul(t16, t16)
	}

	// Step 43: t16 = x^0x172f87d
	t16 = api.Mul(t7, t16)

	// Step 49: t16 = x^0x5cbe1f40
	for s := 0; s < 6; s++ {
		t16 = api.Mul(t16, t16)
	}

	// Step 50: t15 = x^0x5cbe1f75
	t15 = api.Mul(t15, t16)

	// Step 55: t15 = x^0xb97c3eea0
	for s := 0; s < 5; s++ {
		t15 = api.Mul(t15, t15)
	}

	// Step 56: t15 = x^0xb97c3eeb7
	t15 = api.Mul(t14, t15)

	// Step 62: t15 = x^0x2e5f0fbadc0
	for s := 0; s < 6; s++ {
		t15 = api.Mul(t15, t15)
	}

	// Step 63: t14 = x^0x2e5f0fbadd7
	t14 = api.Mul(t14, t15)

	// Step 71: t14 = x^0x2e5f0fbadd700
	for s := 0; s < 8; s++ {
		t14 = api.Mul(t14, t14)
	}

	// Step 72: t13 = x^0x2e5f0fbadd723
	t13 = api.Mul(t13, t14)

	// Step 80: t13 = x^0x2e5f0fbadd72300
	for s := 0; s < 8; s++ {
		t13 = api.Mul(t13, t13)
	}

	// Step 81: t13 = x^0x2e5f0fbadd72321
	t13 = api.Mul(t12, t13)

	// Step 87: t13 = x^0xb97c3eeb75c8c840
	for s := 0; s < 6; s++ {
		t13 = api.Mul(t13, t13)
	}

	// Step 88: t13 = x^0xb97c3eeb75c8c873
	t13 = api.Mul(t0, t13)

	// Step 94: t13 = x^0x2e5f0fbadd72321cc0
	for s := 0; s < 6; s++ {
		t13 = api.Mul(t13, t13)
	}

	// Step 95: t12 = x^0x2e5f0fbadd72321ce1
	t12 = api.Mul(t12, t13)

	// Step 102: t12 = x^0x172f87dd6eb9190e7080
	for s := 0; s < 7; s++ {
		t12 = api.Mul(t12, t12)
	}

	// Step 103: t11 = x^0x172f87dd6eb9190e70a5
	t11 = api.Mul(t11, t12)

	// Step 111: t11 = x^0x172f87dd6eb9190e70a500
	for s := 0; s < 8; s++ {
		t11 = api.Mul(t11, t11)
	}

	// Step 112: t11 = x^0x172f87dd6eb9190e70a52b
	t11 = api.Mul(t10, t11)

	// Step 118: t11 = x^0x5cbe1f75bae46439c294ac0
	for s := 0; s < 6; s++ {
		t11 = api.Mul(t11, t11)
	}

	// Step 119: t11 = x^0x5cbe1f75bae46439c294acd
	t11 = api.Mul(z, t11)

	// Step 127: t11 = x^0x5cbe1f75bae46439c294acd00
	for s := 0; s < 8; s++ {
		t11 = api.Mul(t11, t11)
	}

	// Step 128: t11 = x^0x5cbe1f75bae46439c294acd33
	t11 = api.Mul(t0, t11)

	// Step 134: t11 = x^0x172f87dd6eb9190e70a52b34cc0
	for s := 0; s < 6; s++ {
		t11 = api.Mul(t11, t11)
	}

	// Step 135: t10 = x^0x172f87dd6eb9190e70a52b34ceb
	t10 = api.Mul(t10, t11)

	// Step 141: t10 = x^0x5cbe1f75bae46439c294acd33ac0
	for s := 0; s < 6; s++ {
		t10 = api.Mul(t10, t10)
	}

	// Step 142: t10 = x^0x5cbe1f75bae46439c294acd33ae7
	t10 = api.Mul(t4, t10)

	// Step 145: t10 = x^0x2e5f0fbadd72321ce14a56699d738
	for s := 0; s < 3; s++ {
		t10 = api.Mul(t10, t10)
	}

	// Step 146: t9 = x^0x2e5f0fbadd72321ce14a56699d73f
	t9 = api.Mul(t9, t10)

	// Step 161: t9 = x^0x172f87dd6eb9190e70a52b34ceb9f8000
	for s := 0; s < 15; s++ {
		t9 = api.Mul(t9, t9)
	}

	// Step 162: t8 = x^0x172f87dd6eb9190e70a52b34ceb9f8011
	t8 = api.Mul(t8, t9)

	// Step 167: t8 = x^0x2e5f0fbadd72321ce14a56699d73f00220
	for s := 0; s < 5; s++ {
		t8 = api.Mul(t8, t8)
	}

	// Step 168: t8 = x^0x2e5f0fbadd72321ce14a56699d73f00221
	t8 = api.Mul(x, t8)

	// Step 175: t8 = x^0x172f87dd6eb9190e70a52b34ceb9f8011080
	for s := 0; s < 7; s++ {
		t8 = api.Mul(t8, t8)
	}

	// Step 176: t7 = x^0x172f87dd6eb9190e70a52b34ceb9f80110bd
	t7 = api.Mul(t7, t8)

	// Step 177: t7 = x^0x2e5f0fbadd72321ce14a56699d73f002217a
	t7 = api.Mul(t7, t7)

	// Step 178: t6 = x^0x2e5f0fbadd72321ce14a56699d73f002217f
	t6 = api.Mul(t6, t7)

	// Step 188: t6 = x^0xb97c3eeb75c8c873852959a675cfc00885fc00
	for s := 0; s < 10; s++ {
		t6 = api.Mul(t6, t6)
	}

	// Step 189: t5 = x^0xb97c3eeb75c8c873852959a675cfc00885fc39
	t5 = api.Mul(t5, t6)

	// Step 195: t5 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e40
	for s := 0; s < 6; s++ {
		t5 = api.Mul(t5, t5)
	}

	// Step 196: t4 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e67
	t4 = api.Mul(t4, t5)

	// Step 201: t4 = x^0x5cbe1f75bae46439c294acd33ae7e00442fe1cce0
	for s := 0; s < 5; s++ {
		t4 = api.Mul(t4, t4)
	}

	// Step 202: t4 = x^0x5cbe1f75bae46439c294acd33ae7e00442fe1ccf3
	t4 = api.Mul(t2, t4)

	// Step 210: t4 = x^0x5cbe1f75bae46439c294acd33ae7e00442fe1ccf300
	for s := 0; s < 8; s++ {
		t4 = api.Mul(t4, t4)
	}

	// Step 211: t4 = x^0x5cbe1f75bae46439c294acd33ae7e00442fe1ccf333
	t4 = api.Mul(t0, t4)

	// Step 218: t4 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e6799980
	for s := 0; s < 7; s++ {
		t4 = api.Mul(t4, t4)
	}

	// Step 219: t3 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f
	t3 = api.Mul(t3, t4)

	// Step 228: t3 = x^0x5cbe1f75bae46439c294acd33ae7e00442fe1ccf3331e00
	for s := 0; s < 9; s++ {
		t3 = api.Mul(t3, t3)
	}

	// Step 229: t3 = x^0x5cbe1f75bae46439c294acd33ae7e00442fe1ccf3331e33
	t3 = api.Mul(t0, t3)

	// Step 236: t3 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f1980
	for s := 0; s < 7; s++ {
		t3 = api.Mul(t3, t3)
	}

	// Step 237: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f1993
	t2 = api.Mul(t2, t3)

	// Step 245: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f199300
	for s := 0; s < 8; s++ {
		t2 = api.Mul(t2, t2)
	}

	// Step 246: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f199333
	t2 = api.Mul(t0, t2)

	// Step 254: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f19933300
	for s := 0; s < 8; s++ {
		t2 = api.Mul(t2, t2)
	}

	// Step 255: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f19933333
	t2 = api.Mul(t0, t2)

	// Step 263: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f1993333300
	for s := 0; s < 8; s++ {
		t2 = api.Mul(t2, t2)
	}

	// Step 264: t2 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f1993333333
	t2 = api.Mul(t0, t2)

	// Step 270: t2 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664ccccccc0
	for s := 0; s < 6; s++ {
		t2 = api.Mul(t2, t2)
	}

	// Step 271: t1 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb
	t1 = api.Mul(t1, t2)

	// Step 279: t1 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb00
	for s := 0; s < 8; s++ {
		t1 = api.Mul(t1, t1)
	}

	// Step 280: t1 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb33
	t1 = api.Mul(t0, t1)

	// Step 288: t1 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb3300
	for s := 0; s < 8; s++ {
		t1 = api.Mul(t1, t1)
	}

	// Step 289: t1 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb3333
	t1 = api.Mul(t0, t1)

	// Step 297: t1 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb333300
	for s := 0; s < 8; s++ {
		t1 = api.Mul(t1, t1)
	}

	// Step 298: t0 = x^0xb97c3eeb75c8c873852959a675cfc00885fc399e6663c664cccccccb333333
	t0 = api.Mul(t0, t1)

	// Step 304: t0 = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f19933333332ccccccc0
	for s := 0; s < 6; s++ {
		t0 = api.Mul(t0, t0)
	}

	// Step 305: z = x^0x2e5f0fbadd72321ce14a56699d73f002217f0e679998f19933333332cccccccd
	z = api.Mul(z, t0)

	return z
}

func non_linear(api frontend.API, x0 frontend.Variable, x1 frontend.Variable, x2 frontend.Variable) (frontend.Variable, frontend.Variable, frontend.Variable) {

	alpha, _ := new(big.Int).SetString("20950244155795017333954742965657628047481163604901233004908207073969011285354", 10)
	beta, _ := new(big.Int).SetString("3710185818436319233594998810848289882480745979515096857371562288200759554874", 10)

	y0 := add_chain_exp_d_inv(api, x0)
	y1 := exp_d(api, x1)

	// l for t = 3
	l := api.Add(y0, y1)
	lq := api.Mul(l, l)
	l = api.Mul(l, alpha)
	l = api.Add(l, lq)
	l = api.Add(l, beta)
	y2 := api.Mul(x2, l)

	return y0, y1, y2

}

func mds_t_3_no_rc(api frontend.API, x0 frontend.Variable, x1 frontend.Variable, x2 frontend.Variable) (frontend.Variable, frontend.Variable, frontend.Variable) {

	/* for t = 3, M = circ(2,1,1)
	   for state input (a, b, c) output is =

	   | 2a + b  + c  |
	   | a  + 2b + c  |
	   | a  + b  + 2c |
	*/

	sum := api.Add(x0, x1, x2)
	return api.Add(sum, x0), api.Add(sum, x1), api.Add(sum, x2)
}

func mds_t_3_rc(api frontend.API, x0 frontend.Variable, x1 frontend.Variable, x2 frontend.Variable, rc0 frontend.Variable, rc1 frontend.Variable, rc2 frontend.Variable) (frontend.Variable, frontend.Variable, frontend.Variable) {

	/* for t = 3, M = circ(2,1,1)
	   for state input (a, b, c) output is =

	   | 2a + b  + c  |
	   | a  + 2b + c  |
	   | a  + b  + 2c |
	*/

	sum := api.Add(x0, x1, x2)
	return api.Add(sum, x0, rc0), api.Add(sum, x1, rc1), api.Add(sum, x2, rc2)
}

// Define declares the circuit constraints

func (circuit *Circuit) Define(api frontend.API) error {

	rc0_0, _ := new(big.Int).SetString("4b74133a5743b7fb3b931a7686b7af77c6006da49d8bf85b39de9328ad2b9d87", 16)
	rc0_1, _ := new(big.Int).SetString("0738f78efd46ac801667362968f2ed3e93a588629f0a8067af537055ef3b8de5", 16)
	rc0_2, _ := new(big.Int).SetString("08606cb2fdb16bc68da0ac1809d90f0f3ed1c8cb6ffb202ae68e0483dab296eb", 16)

	rc1_0, _ := new(big.Int).SetString("3bbf8bece610e71898268fa9e82c81c53c8a13d0e5dfe0eac3dc41c01b43c56e", 16)
	rc1_1, _ := new(big.Int).SetString("32df86c26e4cba04adc1eaf7d87ecafe35760975fcf80980598d927fceb2fc8e", 16)
	rc1_2, _ := new(big.Int).SetString("3bdf35c76f89efa67dc23640108db6d25b021cebb054be743dc6346bf9626b53", 16)

	rc2_0, _ := new(big.Int).SetString("671b2f63f8c04c9a19be8f4872289cb815d027c2c91b594e0aeaf173cc3b8bb5", 16)
	rc2_1, _ := new(big.Int).SetString("68346d08afc4f82c3657aeb8999e0f9a16e654619e26a63468a394bcd8f3649d", 16)
	rc2_2, _ := new(big.Int).SetString("153d2fa172d4bc90d53fd5d6f198f5e8b14632a8388d9c76be162b4e9992b833", 16)

	rc3_0, _ := new(big.Int).SetString("28d1f8823d6a9a7532056aafda15c3f5f2ad898833c63a0391a5db57d297d2d4", 16)
	rc3_1, _ := new(big.Int).SetString("3ef2caf6b3356646639487a4eef208fcba88b366518fccd3cbbdd115b324d450", 16)
	rc3_2, _ := new(big.Int).SetString("4df9667b63073058c6de8747008bdc5fc916180c560dc2bd68a1ef5fa28ffc6e", 16)

	rc4_0, _ := new(big.Int).SetString("576a9cbd19be1ba96436f2836af9855527b0e974cd119049ed418d4a85486945", 16)
	rc4_1, _ := new(big.Int).SetString("73ce145e8a5618e247a02a10376451e72be777043568b3cd1239594f5c534e16", 16)
	rc4_2, _ := new(big.Int).SetString("2a609446e01675490a9b3c51f13ee5794a32c6f3d81ce9c9a5897156d7348033", 16)

	rc5_0, _ := new(big.Int).SetString("41b1d71f98a56c8edfb28f3aa869c23c41ff2b5705c54e6958506dd414b22459", 16)
	rc5_1, _ := new(big.Int).SetString("4fe55dab4d90e4d3969fe62909931cb278862a2da4404ff819ab22c7839e47c3", 16)
	rc5_2, _ := new(big.Int).SetString("61245bac99f9b2a6cfd572677c0384e11928a9e11fe47d40d4648fb4885d1bb0", 16)

	rc6_0, _ := new(big.Int).SetString("28968c2074d47eee8c40536ee3cbd82ef412ffbd79ba64c62d0c978a97f5754e", 16)
	rc6_1, _ := new(big.Int).SetString("2c19713b52a1ae3565789a59051530429a26714eb1433e75f5cf6ab9ec71a9d1", 16)
	rc6_2, _ := new(big.Int).SetString("461e6c2ad5ea136098a6a15d012bed4da101c302b588fc2949fd14c99ccae765", 16)

	rc7_0, _ := new(big.Int).SetString("3eedbd6d89cd1fa30b50dbe3b3bd959ff83168c417d1ce5f60e1c2174058a2c5", 16)
	rc7_1, _ := new(big.Int).SetString("2ac8dbbb2704463c04245bf11a272e1fb2e6d283786a23314fc7d27ec1bb10e0", 16)
	rc7_2, _ := new(big.Int).SetString("34c930ef1931d6571eb3d1a37d699aeed1da7209f7366f6fd218235aa8ef3743", 16)

	rc8_0, _ := new(big.Int).SetString("5037e0f7fa3448994db966638aa77092ab5fabec8f32a322fea2a1f2c26f22cc", 16)
	rc8_1, _ := new(big.Int).SetString("480523d6a067ce2a2bee34c7f7e3ff839fb621ba030fb2072f3a8e66b3f3a437", 16)
	rc8_2, _ := new(big.Int).SetString("69dfffdedfef891b448a1adee96de4d8828ac44d39803bc7ac55b62acf948e23", 16)

	rc9_0, _ := new(big.Int).SetString("307f367156f88b9f63ae791d3de97ea7f56152107fcd448ef9741b0d9c6bad19", 16)
	rc9_1, _ := new(big.Int).SetString("0eb8c7dd9af6474636fd3970460a24e5ba90cd0c5b9fd14991791a4610763e51", 16)
	rc9_2, _ := new(big.Int).SetString("37a5ee57a15ef2e4e64661979b98de87d2d9e07b83977fc86ad7d850e9bc48a7", 16)

	rc10_0, _ := new(big.Int).SetString("0886bf20a448bd83746fe5e21995cf73911701f7ee2761575d47c44b3827dde5", 16)
	rc10_1, _ := new(big.Int).SetString("70e9886f1de2e576b86c893ddf531b467cd42ea6eb6e4273034053284ff7cb02", 16)
	rc10_2, _ := new(big.Int).SetString("4f1b3f6a9f91cfc548389964eefe28a3a9572224c5541962011b540f60b509d8", 16)

	s0, s1, s2 := mds_t_3_no_rc(api, circuit.X0, circuit.X1, circuit.X2)

	// round 1
	y0, y1, y2 := non_linear(api, s0, s1, s2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc0_0, rc0_1, rc0_2)

	// round 2
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc1_0, rc1_1, rc1_2)

	// round 3
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc2_0, rc2_1, rc2_2)

	// round 4
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc3_0, rc3_1, rc3_2)

	// round 5
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc4_0, rc4_1, rc4_2)

	// round 6
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc5_0, rc5_1, rc5_2)

	// round 7
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc6_0, rc6_1, rc6_2)

	// round 8
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc7_0, rc7_1, rc7_2)

	// round 9
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc8_0, rc8_1, rc8_2)

	// round 10
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc9_0, rc9_1, rc9_2)

	// round 11
	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_rc(api, y0, y1, y2, rc10_0, rc10_1, rc10_2)

	// round 12

	y0, y1, y2 = non_linear(api, y0, y1, y2)
	y0, y1, y2 = mds_t_3_no_rc(api, y0, y1, y2)

	api.AssertIsEqual(circuit.Y0, y0)
	api.AssertIsEqual(circuit.Y1, y1)
	api.AssertIsEqual(circuit.Y2, y2)

	return nil
}
