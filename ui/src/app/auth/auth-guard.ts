import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';

export const authGuard: CanActivateFn = () => {
  const token = localStorage.getItem('token');
  const router = inject(Router);
  console.log(token);
  if (token) {
    return true;
  } else {
    console.log("Null token!");
    router.navigate(['/login']);
    return false;
  }
};

