import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { tap } from 'rxjs';
import { SampleModuleActionsService } from '../generated/SampleModuleActionsService';
import { MyGetSinglePostRes } from './app.overrides';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, CommonModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent implements OnInit {
  data: MyGetSinglePostRes | null = null;
  constructor(private service: SampleModuleActionsService) {}
  title = 'angular';

  ngOnInit(): void {
    this.service
      .getSinglePost({ id: 1 })
      .pipe(
        tap((res) => {
          this.data = res;
        })
      )
      .subscribe();
  }
}
