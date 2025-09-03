import { Routes } from '@angular/router';
import { Home } from './home/home';
import { Login } from './auth/login/login';

export const routes: Routes = [
  { path: 'login', component: Login },
  { path: 'home', component: Home },
  { path: '', redirectTo: 'login', pathMatch: 'full' }
];

