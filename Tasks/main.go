package main

import ("fmt")

type taskList struct{
	tasks []*task
}

func (t *taskList) addToList(newTask *task){
	// APPEND AGREGA ELEMENTOS A UN SLICE (ARRAY DE LONGITUD VARIABLE)
	t.tasks = append(t.tasks, newTask)
}

func (t *taskList) deleteFromList(index int){
	// PERMITE TOMAR LOS VALORES ANTES DE UN INDICE Y LOS VALORES DESPUÉS DEL INDICE
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:] ...)
}

func (t * taskList) showList(){
	for _,task := range t.tasks{
		fmt.Println("Nombre", task.name )
		fmt.Println("Descripción", task.description )
		fmt.Println("Completada", task.done )
		fmt.Println("------------------------------------------")
	}
}

func (t * taskList) showDoneList(){
	for _,task := range t.tasks{
		if task.done{
		fmt.Println("Nombre", task.name )
		fmt.Println("Descripción", task.description )
		fmt.Println("Completada", task.done )	
		fmt.Println("------------------------------------------")
		}
	}
}

type task struct{
	name string
	description string
	done bool
}

func (t *task) isDone(){
	t.done = true
}

func (t *task) updateDescription(description string){
	t.description = description
}

func (t *task) updateName(name string){
	t.name = name
}

func main(){
	t1 := &task{
		name: "Completar mi curso de Go",
		description: "Completar el curso en Platzi",
	}
	t2 := &task{
		name: "Completar mi curso de Python",
		description: "Completar el curso en Platzi",
	}
	t3 := &task{
		name: "Completar mi curso de Node.js",
		description: "Completar el curso en Platzi",
		done:true,
	}

	// %+v indica que va a imprimir una interfaz de forma att:val
	// fmt.Printf("%+v", t1);
	// len DEVUELVE LA LONGITUD
	// fmt.Println(len(list.tasks))


	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	list.addToList(t3);
	
	// for i:=0; i< 10; i++{
	// 	if i==5{
	// 		// CONTINUE ROMPE TODO EL CICLO
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("------------------------------------------")
	// for i:=0; i< 10; i++{
	// 	if i==5{
	// 		// CONTINUE ROMPE LA ITERACIÓN PERO SIGUE CON EL CICLO
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("------------------------------------------")
	list.showList()
	list.showDoneList()
	
	// CREAMOS UN MAPA PARA GUARDAR DIFERENTES LISTAS DE DIFERENTES PERSONAS
	tasksMap := make(map[string]*taskList)
	tasksMap["Andoni"] = list

	t4 := &task{
		name: "Completar mi curso de GIT",
		description: "Completar el curso en Platzi",
	}
	t5 := &task{
		name: "Completar mi curso de Golang",
		description: "Completar el curso en Platzi",
		done:true,
	}
	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	tasksMap["Moni"] = list2
	fmt.Println("TAREAS DE ANDONI")
	tasksMap["Andoni"].showList()
	fmt.Println("TAREAS DE MONI")
	tasksMap["Moni"].showList()

}