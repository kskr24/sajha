import { Routes } from '@angular/router';
import { Hello } from './hello/hello';
import { Login } from './auth/login/login';
import { authGuard } from './auth/auth-guard';
import { Board } from './board/board';

export const routes: Routes = [
  {path: 'boards', component: Board},
  { path: 'login', component: Login },
  { path: 'hello', component: Hello, canActivate: [authGuard] },
  { path: '', redirectTo: 'login', pathMatch: 'full' }
];

