package library

type Watcher int

func (w *Watcher) GetInfo(arg int,result *int) error{
    *result = 1
    return nil
}
