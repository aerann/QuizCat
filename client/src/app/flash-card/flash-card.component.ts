import { Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-flash-card',
  templateUrl: './flash-card.component.html',
  styleUrl: './flash-card.component.css'
})
export class FlashCardComponent implements OnInit{
  @Input() cardData: any; 
  isClicked;

  ngOnInit(){
    // cardData: any
    console.log(this.cardData)
  }
  
  toggleClicked(){
    this.isClicked = !this.isClicked 
    return this.isClicked
  }

  getColor(){
    if (!this.isClicked){
      return '#a3c3eb' //change to background of card
    } 
    return '#afe1a6'
  }
}
