import { Component, OnInit } from '@angular/core';
import { Auth } from '../auth/auth';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hello',
  imports: [],
  templateUrl: './hello.html',
  styleUrl: './hello.css'
})
export class Hello implements OnInit {
  message = '';
  constructor(private auth: Auth, private route: Router) {}
  ngOnInit() {
    this.auth.getHello().subscribe({
      next: (res) => {
        this.message = res.message;
      },
      error: () => {
        this.message = 'Unauthorized';
      }
    });
  }

  logout() {
    localStorage.removeItem('token');
    this.route.navigate(['/login']);
  }
}

