import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-flash-card',
  templateUrl: './flash-card.component.html',
  styleUrl: './flash-card.component.css',
})
export class FlashCardComponent implements OnInit {
  @Input() cardData: any;
  isClicked = false;

  ngOnInit() {
    // cardData: any
    console.log(this.cardData);
  }

  toggleClicked() {
    this.isClicked = !this.isClicked;
    // return this.isClicked;
  }

  getColor() {
    if (!this.isClicked) {
      return '#fbb37ca2'; //change to background of card
    }
    return '#afe1a6';
  }
}
