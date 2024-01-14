import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { StudyMaterialComponent } from './study-material/study-material.component';
import { HttpClientModule } from '@angular/common/http';
import { FlashCardComponent } from './flash-card/flash-card.component';

@NgModule({
  declarations: [
    AppComponent,
    StudyMaterialComponent,
    FlashCardComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule 
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
