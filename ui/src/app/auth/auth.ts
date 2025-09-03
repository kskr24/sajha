import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class Auth {
  private baseURL = 'http://localhost:8080/v1';
  private loggedIn = new BehaviorSubject<boolean>(false);

  constructor(private http: HttpClient) {}
  login(username: string, password: string): Observable<any> {
    return this.http.post(`${this.baseURL}/login`, { username, password });
  }
  
  getHello(): Observable<any> {
    return this.http.get(`${this.baseURL}/hello`, {
      headers: { Authorization: localStorage.getItem('token') || '' }
    });
  }
  setLoggedIn(state: boolean) {
    return this.loggedIn.asObservable();
  }
  logout() {
    localStorage.removeItem('token');
    this.loggedIn.next(false);
  }
}

