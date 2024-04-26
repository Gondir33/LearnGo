package main

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedList_Load(t *testing.T) {
	d := &DoubleLinkedList{}
	d.LoadData("test.json")
	assert.NotNil(t, d.curr)

	d1 := &DoubleLinkedList{}
	err := d1.LoadData("sadsads")
	assert.NotNil(t, err)

	err = d1.LoadData("testerr")
	assert.NotNil(t, err)
}

func TestDoubleLinkedList_Operations(t *testing.T) {
	// Создаем экземпляр DoubleLinkedList и инициализируем его тестовыми данными
	var d DoubleLinkedList
	testData := GenerateData()
	d.Init(testData)

	// Проверяем операцию Len
	assert.Equal(t, len(testData), d.Len())

	// Проверяем операцию Current
	assert.NotNil(t, d.Current())

	// Проверяем операции Next и Prev
	initialNode := d.Current()
	nextNode := d.Next()
	assert.NotNil(t, nextNode)
	assert.Equal(t, initialNode, nextNode.prev)

	initialNode = d.Current()
	prevNode := d.Prev()
	assert.NotNil(t, prevNode)
	assert.Equal(t, initialNode, prevNode.next)

	// Проверяем операцию Insert
	newCommit := Commit{
		Message: "New Commit",
		UUID:    gofakeit.UUID(),
		Date:    gofakeit.Date(),
	}
	err := d.Insert(0, newCommit)
	assert.NoError(t, err)
	assert.Equal(t, len(testData)+1, d.Len())
	tmpNode, err := d.GetByIndex(1)
	assert.Equal(t, newCommit.Message, tmpNode.data.Message)

	err = d.Insert(d.len-1, newCommit)
	assert.Equal(t, len(testData)+2, d.Len())
	tmpNode, err = d.GetByIndex(d.len - 1)
	assert.Equal(t, newCommit.Message, tmpNode.data.Message)

	// Проверяем операции Delete и DeleteCurrent
	err = d.Delete(1)
	err = d.Delete(d.len - 1)
	assert.NoError(t, err)
	assert.Equal(t, len(testData), d.Len())

	err = d.DeleteCurrent()
	assert.NoError(t, err)
	assert.Equal(t, len(testData)-1, d.Len())

	// Проверяем операцию Index
	index, err := d.Index()
	assert.NoError(t, err)
	assert.Equal(t, 0, index)

	// Проверяем операции Pop и Shift
	poppedNode := d.Pop()
	shiftedNode := d.Shift()
	assert.NotNil(t, poppedNode)
	assert.NotNil(t, shiftedNode)
	assert.Equal(t, shiftedNode, d.head)
	assert.Equal(t, poppedNode, d.tail)

	// Проверяем операции SearchUUID и Search
	uuidToSearch := testData[4].UUID
	messageToSearch := testData[4].Message
	nodeByUUID := d.SearchUUID(uuidToSearch)
	nodeByMessage := d.Search(messageToSearch)
	assert.NotNil(t, nodeByUUID)
	assert.NotNil(t, nodeByMessage)
	assert.Equal(t, uuidToSearch, nodeByUUID.data.UUID)
	assert.Equal(t, messageToSearch, nodeByMessage.data.Message)

	d.SetCurrent(d.len - 1)
	assert.Equal(t, d.tail, d.curr)
	index, err = d.Index()
	assert.Equal(t, index, d.len-1)

	// Проверяем операцию Reverse
	reversedList := d.Reverse()
	assert.NotNil(t, reversedList)
	assert.Equal(t, len(testData)-1, reversedList.Len())
}

func TestDoubleLinkedList_Operations_EdgesCases_Errors(t *testing.T) {
	d := &DoubleLinkedList{}

	err := d.DeleteCurrent()
	assert.NotNil(t, err)

	_, err = d.Index()
	assert.NotNil(t, err)

	testData := GenerateData()
	d.Init(testData)

	err = d.Insert(100, Commit{Message: "errorpls"})
	assert.NotNil(t, err)
	err = d.Delete(100)
	assert.NotNil(t, err)

	wantAfterDelete := d.curr.next
	d.Delete(0)
	assert.Equal(t, wantAfterDelete, d.curr)

	wantAfterDelete = d.tail.prev
	d.Delete(d.len - 1)
	assert.Equal(t, wantAfterDelete, d.tail)

	d.SetCurrent(d.len - 1)
	wantAfterDelete = d.tail.prev
	d.DeleteCurrent()
	assert.Equal(t, wantAfterDelete, d.tail)

	err = d.SetCurrent(100)
	assert.NotNil(t, err)

	_, err = d.GetByIndex(100)
	assert.NotNil(t, err)

	node := d.SearchUUID("asd")
	assert.Nil(t, node)

	node = d.Search("asd")
	assert.Nil(t, node)

	newCommit := Commit{
		Message: "New Commit",
		UUID:    gofakeit.UUID(),
		Date:    gofakeit.Date(),
	}
	d.Push(newCommit)
	assert.Equal(t, *d.tail.data, newCommit)

	d.Delete(4)
	assert.Equal(t, len(testData)-3, d.len)

	d.SetCurrent(4)
	d.DeleteCurrent()
	assert.Equal(t, len(testData)-4, d.len)

	d.SetCurrent(d.len - 1)
	d.Delete(d.len - 1)
	assert.Equal(t, len(testData)-5, d.len)

}
