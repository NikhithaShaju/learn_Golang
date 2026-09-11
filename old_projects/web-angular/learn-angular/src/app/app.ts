import { Component, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterOutlet } from '@angular/router';

@Component({
  imports: [RouterOutlet, FormsModule],
  selector: 'app-root',
  styleUrl: './app.css',
  templateUrl: './app.html',
})
export class App {
  task="";
  tasks: { name: string, completed: boolean }[] = [];
  
 
addTask() {
  if (this.task.trim() === "") {
    return;
  }
  this.tasks.push({
    name: this.task,
    completed: false
  });

  this.task = "";
}
   

deleteTask(item: { name: string, completed: boolean }) {
  const index = this.tasks.indexOf(item);
  this.tasks.splice(index, 1);
}

toggleTask(item: { name: string, completed: boolean }) {
  item.completed = !item.completed;
}
  protected readonly title = signal('learn-angular');
}
