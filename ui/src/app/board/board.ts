import {
  CdkDragDrop,
  moveItemInArray,
  transferArrayItem,
  CdkDropList
} from '@angular/cdk/drag-drop';
import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';

@Component({
  selector: 'app-board',
  imports: [CommonModule, CdkDropList],
  templateUrl: './board.html',
  styleUrl: './board.css'
})
export class Board {
  columns = [
    { name: 'Todo', tasks: ['Task A', 'Task B', 'Task C'] },
    { name: 'In Progress', tasks: ['Task D'] },
    { name: 'Done', tasks: ['Task E'] }
  ];

  drop(event: CdkDragDrop<string[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(
        event.container.data,
        event.previousIndex,
        event.currentIndex
      );
    } else {
      transferArrayItem(
        event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex
      );
    }
  }
}

