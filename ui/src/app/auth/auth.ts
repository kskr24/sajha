import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class Auth {
  private baseURL = 'http://localhost:42';
  private loggedIn = new BehaviorSubject<boolean>(false);

  constructor(private http: HttpClient) {}
  login(username: string, password: string): Observable<any> {
    //return this.http.post(`${this.baseURL}/login`, { username, password });
    return of({ success: true, token: 'F You' });
  }

  getHello(): Observable<any> {
    // return this.http.get(`${this.baseURL}/hello`, {
    //   headers: { Authorization: localStorage.getItem('token') || '' }
    // });
    return of({ success: true, token: 'F You' });
  }
  setLoggedIn(state: boolean) {
    return this.loggedIn.asObservable();
  }
  logout() {
    localStorage.removeItem('token');
    this.loggedIn.next(false);
  }
}

