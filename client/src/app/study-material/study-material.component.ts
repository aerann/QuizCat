import { Component, OnInit } from '@angular/core';
import { FormsModule, NgForm } from '@angular/forms';
import { HttpClient} from '@angular/common/http';

@Component({
  selector: 'app-study-material',
  templateUrl: './study-material.component.html',
  styleUrl: './study-material.component.css'
})

export class StudyMaterialComponent {
  study_material_text; 
  res;
  isLoading = false; 
  // cards;
  //cards = ['test1', 'test2'] 
  cards = [] 
  constructor(private http: HttpClient) { }


  async onSubmit(form: NgForm){
    
    console.log(form);
    console.log('this is the body text:', this.study_material_text )
    // Simple POST request with a JSON body and response type <any>
    this.isLoading = true
    this.http.post<any>('http://localhost:8080/generate_cards', { text: this.study_material_text}).subscribe(data => {
        this.res = data
        this.isLoading = false;  
        for (const card of this.res.cards) {
          this.cards.push(card)
        }
    })
  }

}
