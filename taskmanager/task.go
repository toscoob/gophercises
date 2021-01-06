package taskmanager

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

var tasksBucket = []byte("tasks")

func itob(v uint64) []byte {
	s := strconv.FormatUint(v, 10)
	return []byte(s)
}

//func btoi(b []byte) uint64 {
//	s, err := strconv.ParseUint(string(b), 10, 64)
//	if err != nil {
//		log.Println(err)
//	}
//	return s
//}

func args2ints(args []string) []int {
	var resp []int

	for _, arg := range args {
		intArg, err := strconv.Atoi(arg)
		if err == nil {
			resp = append(resp, intArg)
		}
	}

	return resp
}

func dbFile() string {
	file := "data.db"
	if dir, err := homedir.Dir(); err == nil {
		file = filepath.Join(dir, file)
	}

	return file
}

//todo set complete date instead of deleting tasks
//--- ADD
type AddCommand struct{}

func (t *AddCommand) Help() string {
	return "add [arg0]"
}

func (t *AddCommand) Run(args []string) int {
	//fmt.Println("add", args)
	if len(args) < 1 {
		log.Fatal("please provide task description")
		return 1
	}

	db, err := bolt.Open(dbFile(), 0600, nil)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	defer db.Close()

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return err
		}

		id, _ := bucket.NextSequence()
		buf := strings.Join(args[:]," ")

		err = bucket.Put(itob(id), []byte(buf))
		if err != nil {
			return err
		}
		fmt.Printf("Added \"%s\" to your task list.\n", buf)
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}

func (t *AddCommand) Synopsis() string {
	return "Add a new task to your TODO list"
}
//--- DO
type DoCommand struct{}

func (t *DoCommand) Help() string {
	return "do [arg0] [arg1] "
}

func (t *DoCommand) Run(args []string) int {
	//fmt.Println("do", args)

	if len(args) < 1 {
		log.Fatal("please provide task id(s)")
		return 1
	}

	intArgs := args2ints(args)
	if len(intArgs) < 1 {
		log.Fatal("please provide integer task id(s)")
		return 1
	}

	db, err := bolt.Open(dbFile(), 0600, nil)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	defer db.Close()

	// try to read
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		if b == nil {
			return errors.New("bucket does not exist")
		}

		//buf := bucket.Get([]byte(args[0]))
		//if buf == nil {
		//	return errors.New(fmt.Sprintf("task %s not found", args[0]))
		//}
		////todo change to mark complete later
		//err = bucket.Delete([]byte(args[0]))
		//
		//if err != nil {
		//	return err
		//}

		var buf []byte
		c := b.Cursor()
		idx := 1 //not sure how to assign in loop statement together with k, v
		deleted := false
		for k, v := c.First(); k != nil; k, v = c.Next() {
			//fmt.Printf("%d. %s\n", idx, v)
			if idx == intArgs[0] { //todo remake to handle multiple ids
				err = b.Delete(k)
				if err != nil {
					return err
				}
				buf = v
				deleted = true
				break
			}
			idx += 1
		}

		if !deleted {
			return errors.New(fmt.Sprintf("task %d not found", intArgs[0]))
		}

		fmt.Printf("You have completed the \"%s\" task\n", buf)
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}

func (t *DoCommand) Synopsis() string {
	return "Mark a task on your TODO list as complete"
}
//--- LIST
type ListCommand struct{}

func (t *ListCommand) Help() string {
	return "add [arg0] [arg1] "
}

func (t *ListCommand) Run(args []string) int {
	//fmt.Println("This is a fake \"list\" command")
	db, err := bolt.Open(dbFile(), 0600, nil)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		if b == nil {
			return errors.New("bucket does not exist")
		}

		fmt.Println("You have the following tasks:")
		c := b.Cursor()
		idx := 1 //not sure how to assign in loop statement together with k, v
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%d. %s\n", idx, v)
			idx += 1
		}
		//err = b.ForEach(func(k, v []byte) error {
		//	fmt.Printf("%d. %s\n", btoi(k), v)
		//	return nil
		//})
		//
		//if err != nil {
		//	return err
		//}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		return 1
	}

	return 0
}

func (t *ListCommand) Synopsis() string {
	return "List all of your incomplete tasks"
}

//--- RESET
type ResetCommand struct{}

func (t *ResetCommand) Help() string {
	return "reset [arg0] [arg1] "
}

func (t *ResetCommand) Run(args []string) int {
	fmt.Println("all data is deleted")
	db, err := bolt.Open(dbFile(), 0600, nil)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(tasksBucket)
		if err != nil {
			fmt.Println(err)
			fmt.Println("doesn't matter anyway")
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return 1
	}

	return 0
}

func (t *ResetCommand) Synopsis() string {
	return "Remove all tasks"
}
