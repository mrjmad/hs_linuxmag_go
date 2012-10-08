package addhsgo

import "testing"

func TestAdd(t *testing.T) {
    r := Add(41,1)
    if r != 42 {
        t.Errorf("Add de 41 et 1 donne %v au lieu de 42\n",r)
    }
}

func TestInvalidAddwithError(t *testing.T) {
    r := InvalidAdd(41,1)
    if r != 42 {
        t.Errorf("InvalidAdd de 41 et 1 donne %v au lieu de 42\n",r)
    }
    r = InvalidAdd(32,1)
    if r != 33 {
        t.Errorf("InvalidAdd de 32 et 1 donne %v au lieu de 33\n",r)
    }

}

func TestAddFail(t *testing.T) {
    r := Add(41,1)
    t.Fail()
    if r != 42 {
        t.Errorf("InvalidAdd de 41 et 1 donne %v au lieu de 42\n",r)
    }

    t.Log("On a lancé un fail mais on est toujours dans la fonction TestAddFail")
}

func TestAddFailNow(t *testing.T) {
    r := Add(41,1)
    t.FailNow()
    t.Log("On a lancé un failnow jamais on ne passera ici")
    if r != 42 {
        t.Errorf("InvalidAdd de 41 et 1 donne %v au lieu de 42\n",r)
    }

}
