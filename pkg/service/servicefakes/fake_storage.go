// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	context "context"
	sync "sync"

	service "github.com/penkovski/ginkgo-example/pkg/service"
)

type FakeStorage struct {
	FileStub        func(context.Context, string, string) (*service.File, error)
	fileMutex       sync.RWMutex
	fileArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	fileReturns struct {
		result1 *service.File
		result2 error
	}
	fileReturnsOnCall map[int]struct {
		result1 *service.File
		result2 error
	}
	SaveFileStub        func(context.Context, *service.File, string) error
	saveFileMutex       sync.RWMutex
	saveFileArgsForCall []struct {
		arg1 context.Context
		arg2 *service.File
		arg3 string
	}
	saveFileReturns struct {
		result1 error
	}
	saveFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) File(arg1 context.Context, arg2 string, arg3 string) (*service.File, error) {
	fake.fileMutex.Lock()
	ret, specificReturn := fake.fileReturnsOnCall[len(fake.fileArgsForCall)]
	fake.fileArgsForCall = append(fake.fileArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("File", []interface{}{arg1, arg2, arg3})
	fake.fileMutex.Unlock()
	if fake.FileStub != nil {
		return fake.FileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorage) FileCallCount() int {
	fake.fileMutex.RLock()
	defer fake.fileMutex.RUnlock()
	return len(fake.fileArgsForCall)
}

func (fake *FakeStorage) FileCalls(stub func(context.Context, string, string) (*service.File, error)) {
	fake.fileMutex.Lock()
	defer fake.fileMutex.Unlock()
	fake.FileStub = stub
}

func (fake *FakeStorage) FileArgsForCall(i int) (context.Context, string, string) {
	fake.fileMutex.RLock()
	defer fake.fileMutex.RUnlock()
	argsForCall := fake.fileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStorage) FileReturns(result1 *service.File, result2 error) {
	fake.fileMutex.Lock()
	defer fake.fileMutex.Unlock()
	fake.FileStub = nil
	fake.fileReturns = struct {
		result1 *service.File
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) FileReturnsOnCall(i int, result1 *service.File, result2 error) {
	fake.fileMutex.Lock()
	defer fake.fileMutex.Unlock()
	fake.FileStub = nil
	if fake.fileReturnsOnCall == nil {
		fake.fileReturnsOnCall = make(map[int]struct {
			result1 *service.File
			result2 error
		})
	}
	fake.fileReturnsOnCall[i] = struct {
		result1 *service.File
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) SaveFile(arg1 context.Context, arg2 *service.File, arg3 string) error {
	fake.saveFileMutex.Lock()
	ret, specificReturn := fake.saveFileReturnsOnCall[len(fake.saveFileArgsForCall)]
	fake.saveFileArgsForCall = append(fake.saveFileArgsForCall, struct {
		arg1 context.Context
		arg2 *service.File
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("SaveFile", []interface{}{arg1, arg2, arg3})
	fake.saveFileMutex.Unlock()
	if fake.SaveFileStub != nil {
		return fake.SaveFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveFileReturns
	return fakeReturns.result1
}

func (fake *FakeStorage) SaveFileCallCount() int {
	fake.saveFileMutex.RLock()
	defer fake.saveFileMutex.RUnlock()
	return len(fake.saveFileArgsForCall)
}

func (fake *FakeStorage) SaveFileCalls(stub func(context.Context, *service.File, string) error) {
	fake.saveFileMutex.Lock()
	defer fake.saveFileMutex.Unlock()
	fake.SaveFileStub = stub
}

func (fake *FakeStorage) SaveFileArgsForCall(i int) (context.Context, *service.File, string) {
	fake.saveFileMutex.RLock()
	defer fake.saveFileMutex.RUnlock()
	argsForCall := fake.saveFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStorage) SaveFileReturns(result1 error) {
	fake.saveFileMutex.Lock()
	defer fake.saveFileMutex.Unlock()
	fake.SaveFileStub = nil
	fake.saveFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) SaveFileReturnsOnCall(i int, result1 error) {
	fake.saveFileMutex.Lock()
	defer fake.saveFileMutex.Unlock()
	fake.SaveFileStub = nil
	if fake.saveFileReturnsOnCall == nil {
		fake.saveFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileMutex.RLock()
	defer fake.fileMutex.RUnlock()
	fake.saveFileMutex.RLock()
	defer fake.saveFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ service.Storage = new(FakeStorage)