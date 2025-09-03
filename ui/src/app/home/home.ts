import { Component, OnInit } from '@angular/core';
import { Auth } from '../auth/auth';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  imports: [],
  templateUrl: './home.html',
  styleUrl: './home.css'
})
export class Home implements OnInit {
  message = 'Hello';
  constructor(private auth: Auth, private router: Router) {}

  ngOnInit(): void {
    // const token = localStorage.getItem('token');
    // if(!token){
    //   this.router.navigate(['/login']);
    //   return;
    // }
    // this.auth.getHello(token).subscribe(
    //   {
    //     next: (res: any)=>this.message = res.message,
    //     error: () => this.router.navigate(['/login'])
    //   }
    // );
  }
  logout() {
    this.auth.logout();
    this.router.navigate(['/login']);
  }
}

