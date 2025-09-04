import { ChangeDetectorRef, Component, OnInit, signal } from '@angular/core';
import { Auth } from '../auth/auth';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hello',
  imports: [],
  templateUrl: './hello.html',
  styleUrl: './hello.css'
})
export class Hello implements OnInit {
  message = signal('');
  constructor(private auth: Auth, private route: Router, private cd: ChangeDetectorRef) {}
  ngOnInit() {
    this.auth.getHello().subscribe({
      next: (res) => this.message.set(res.message),
      error: () => this.message.set('Unauthorized')
    });
  }

  logout() {
    localStorage.removeItem('token');
    this.route.navigate(['/login']);
  }
}

