import { Component, OnInit } from '@angular/core';
import { FormsModule, NgForm } from '@angular/forms';
import { HttpClient} from '@angular/common/http';

@Component({
  selector: 'app-study-material',
  templateUrl: './study-material.component.html',
  styleUrl: './study-material.component.css'
})
export class StudyMaterialComponent implements OnInit {
  constructor(private http: HttpClient) { }
  study_material_text; 
  postId;

  ngOnInit() {      
    // Simple POST request with a JSON body and response type <any>
    this.http.post<any>('https://reqres.in/api/posts', { text: this.study_material_text}).subscribe(data => {
        this.postId = data.id; //postId is assigned to  the id from the response? 
    })
  }

  onSubmit(form: NgForm){
    console.log(form);
    console.log('this is the body text:', this.study_material_text )
  }

}
