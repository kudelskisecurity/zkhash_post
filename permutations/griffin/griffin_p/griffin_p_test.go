package griffin

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func TestGriffinPermutation(t *testing.T) {

	var griffinPCircuit Circuit

	// generate R1CS
	r1cs, err := frontend.Compile(ecc.BLS12_381, r1cs.NewBuilder, &griffinPCircuit)

	if err != nil {
		// invalid proof
		t.Log("Invalid r1cs")
	} else {
		t.Log("Valid r1cs")
	}

	// generate proving and verifying keys (groth16 setup)
	pk, vk, err := groth16.Setup(r1cs)

	if err != nil {
		// invalid proof
		t.Log("Invalid setup")
	} else {
		t.Log("Valid setup")
	}

	// generate witness
	griffinPCircuit.X0 = "4508222636468352024383260454905711264828514403328072469752654409237195968161"
	griffinPCircuit.X1 = "23155693710573383855262674279011399341590136940666102282456780392188583899928"
	griffinPCircuit.X2 = "36735696066136083014492339476499380060050269383990014726989163817202767970550"

	griffinPCircuit.Y0 = "31016929337754621893814503718917517227641083441856681906706457834546054116907"
	griffinPCircuit.Y1 = "20864644307131271666118544493239242610879346549161198036990264813404986792643"
	griffinPCircuit.Y2 = "1410085955641781374925541439456535442329813473262597984478561809355127215349"

	witness, err := frontend.NewWitness(&griffinPCircuit, ecc.BLS12_381)

	if err != nil {
		// invalid proof
		t.Log("Invalid witness generation")
	} else {
		t.Log("Valid witness generation")
	}

	publicWitness, err := witness.Public()

	if err != nil {
		// invalid proof
		t.Log("Invalid public witness generation")
	} else {
		t.Log("Valid public witness generation")
	}

	// generate proof
	proof, err := groth16.Prove(r1cs, pk, witness)

	if err != nil {
		// invalid proof
		t.Log("Invalid proof generation")
	} else {
		t.Log("Valid proof generation")
	}

	// verify proof
	err = groth16.Verify(proof, vk, publicWitness)

	if err != nil {
		// invalid proof
		t.Log("Invalid proof verification")
	} else {
		t.Log("Valid proof verification")
	}
}
