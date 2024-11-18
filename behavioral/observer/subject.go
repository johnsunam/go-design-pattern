package main

type Observable interface {
	registerObserver(obs observer)
	unRegisterObserver(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) unRegisterObserver(o DataListener) {
	var newObs []DataListener
	for _, ob := range newObs {
		if ob.Name != o.Name {
			newObs = append(newObs, ob)
		}
	}
	ds.observers = newObs
}

func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onUpdate(ds.field)
	}
}
