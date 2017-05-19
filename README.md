# is105-ica02


### Oppgave 2 
https://github.com/GB-Noname/is105-ica02

hello.go 

### Oppgave 3 
https://github.com/GB-Noname/is105-ica02

sum_test.go 

main_sum.go

### Oppgave 4
https://github.com/GB-Noname/is105-ica02/tree/master/sort

sorting.go

sorting_test.go

### Oppgave 5 
https://github.com/GB-Noname/is105-ica02/tree/master/boring

boring.go

https://github.com/GB-Noname/is105-ica02

boring_main.go


main_boring.goroutine.go


Kode-kommentarer:

Hello.go
I denne filen ligger det en func som heter main. Denne funksjonen gj�r en print av teksten i funksjonen. 

Sum.go
Denne filen inneholder fem funksjoner. Alle returnerer a+b, som er summen av tallene som blir skrevet inn i terminalen. 
Det er fem ulike funksjoner fordi programmet skal kunne opperere med forskjellige typer, og de er som f�lger:int8,
float64, uint32, int32 og int64. 
Int8 var ikke med i oppgaven, men vi tok dette med fordi vi ville bruke den i testen for � f� bedre forst�else for testene, 
samt vi ville kunne forst� rangen bedre med lave tall f�r vi gikk videre til de st�rre tallene. 

Main_sum.go
Her finner vi mainen til sum.go. Denne har en func som heter main, som kj�rer programmet. 
I starten av funksjonen har vi variablen "finish", som blir brukt som en boolean i en for-loop. Denne loopen holder programmet kj�rende,
helt til den brytes, og programmet avsluttes. Etter det blir det deklarert en variabel, typeVal, som er en string. Den bestemmer hvilken 
type som skal kalkuleres. S� kommer det en to prints, som sier hvilken typer som er valide, og at man skal skrive hvilken type man vil kalkulere.
Det man skriver i terminalen blir s� scannet, og den leter etter typeVal.
Videre kommer det fire if statements, hvor koden er lik, forskjellen er bare hvilken type det dreier seg om, for eksemple uint32 eller int32. 
De tar den riktige typen, utifra hva som ble skrevet i terminalen,scanner tallene som blir skrevet inn, s� blir metoden som h�rer til typen kalt 
p� fra sum.go, som legger sammen de to tallene. Det blir brukt to variabler, val og val2. 
N�r programmet har lagt sammen to tall kommer det en print som forklarer hvordan man avslutter programmet. Hvis man skriver "yes", blir finnish==true og
for loopen brytes, og programmet avsluttes. Hvis man skriver "no" kan man fortsette med � legge sammen tall. 

sum_test.go
Et testprogram som tester at kalkulatoren "sum" fungerer som den skal.
Programmet tester med de fem ulike typene. En funksjon for hver type, med en tilh�rende struct. 
I structen er det tre variabler, to for tallene som blir lagt sammen, og �n for tallet som er forventet � f� som svar. 
Tallene som blir testet st�r ogs� i structen. 
I funksjonene er det en for loop med variablen "v", som bestemmer hvilken range som blir brukt, og rangen er lik som tallene i structen. 
S� kommer en if statement som sier at hvis for eksempel if val:= SumInt8 ikke er lik det forventede tallet, printer programmet en error statement, 
som sier hva som ble kalkulert, og hva som var forventet. 

boring.go
Programmet inneholder to funksjoner, Boring01 og Boring10.
Boring01 har "msg" som parameter. Boring01 har en for loop som ikke har noen break. Variablen i begynner p� null, og s� teller den oppover. 
Funksjonen printen en melding eller "msg", og variablen i som stiger med ett nummer for hver linje som blir printet. 
Time.sleep gj�r at den bruker litt tid p� � printe ut neste linje. 
Den andre funcen, Boring10, som har "msg" og "c" som parametere. 
Denne funcen har en lik for loop som Boring01.
s� kommer det en fmt.sprintf som konverterer meldingen til en string fra hvilken som helst type.
Time.sleep gj�r at den bruker litt tid p� � printe ut neste linje. 
C Channel venter p� svar fra sprintf. 



Boring_main.go
Dette er den ene mainen til boring.go. 
Dette programmet kj�rer boring.go, og funksjonen boring01 blir kalt p� fra boring.go. 

main_boring_goroutine.go
Dette er den andre mainen til boring.go. 
Funksjonen boring10 blir kalt p� fra boring.go. 
En channel blir laget og har variabelen c. 
Det er en for loop som som teller oppover fra 0, som er variablen "i".
En melding blir printet.
Boring10 kombinerer teksten som vi sender inn med integeren i loopen p� boring10.  


sorting.go
Et program som sorterer integers i stigende rekkef�lge. 
func bubble_sort_modified har en liste av integers som parameter. 
Variabelen "n" er lik lengde p� listen. 
Variabelen swapped == true, en boolean som er sann n�r integersene er sortert i riktig rekkef�lge. 
En for-loop som sier at n�r swapped==false skal den iterere gjennom tallene, variablen "i" = 1.
N�r "i" er mindre enn lengden p� listen minus 1, skal loopen fortsette. 
If listen [i-1] er st�rre enn listen [i], blir verdiene i tallrekken blir sortert. 
Hvis list[i-1] = list[i-1] er swapped true, og for loopen brytes. 

func bubble_sort har ogs� en liste av integers som parameter. 
Variabelen n = lengden p� listen. 
En for-loop med variabelen i:=0, n�r "i" er mindre enn "n"(lengden p� listen), fortsetter "i" videre p� neste tall. 
En for-loop med variabelen j:=0, n�r "j" er mindre enn "n"(lengden p� lisetn) minus 1, fortsetter "j" videre p� neste tall. 
if list[j] er st�rre enn list [j+1] blir temp:= list[j+1]. 
list [j+1] er da lik list[j]. 
Og det betyr at list[j] er lik temp. 

func Qsort at values som er integers i parameter. 
Denne funksjonen kaller p� qsort og har parameterene values, l:0 og lendgden p� listen som er values -1.

func qsort har tre parametere, values av typen int, l av typen int, og r av typen int. 
if "l" er st�rre eller lik "r", returnerer den en blank return hvis if statementen stemmer. 
Variablen pivot:= values[l]. 
variabelen i:= "l" pluss 1. 
En for-loop med variabelen j:=l. Hvis "j" er mindre eller lik "r" g�r j videre p� neste tall. 
if pivot er st�rre enn values[j], og values[i], values[j] er lik values [j], values [i], g�r i videre p� neste tall. 
Den kj�rer en loop til loopen er komplett, og n�r den er komplett returnerer den blankt. 

sorting_test.go
Et program som kj�rer benchmark tester og test p� funksjoner fra sorting.go.  

 
func init setter den tiden til systemet til seed, og den sikrer at resultatet blir likt hver gang. 

func perm gj�r at en f�r tilfeldige tall � teste med. 

func BenchmarkBsort100 har b*testing.B som parameter, som gj�r det mulig � kj�re benchmark tester. 
Den kaller p� benchmarkBsort funcen som har en int og b som parameter. integeren "i", er hvor mange tall som er listen
som skal sorteres, i dette tilfellet 100. 
Denne funcen gjentar seg flere ganger, bare at variabelen "i" endres til 1000, og 10 000. 

func benchmarkBSort har parameterene "i" og "b". "i" er hvor mange tall som er listen som skal sorteres, og "b" kommer
fra testing packagen som er importert. 
Det er en for-loop med variabelen j:=0, n�r j er mindre enn b.N fortsetter j til neste tall. 
b.StopTimer gj�r at testen blir stoppet.
b.startTimer starter testen igjen. 
Dette er p� grunn av tiden det tar � hente verdiene som skal testes ikke skal bli tatt med i testn

func benchmarkBSortModified er lik som benchmarkBSort, bare at den er for to ulike funksjoner i sorting.go.  

func TestQSort tester QSort funksjonen fra sorting.go. 
Den har variablene values og expected, og begge holder en liste med tall. 
values er ikke sortert, mens expected er tallene sortert i stigende rekkef�lge. 
Hvis QSort ikke har sortert de i riktig rekkef�lge blir det printet en feilmelding. 

func BenchmarkQSort100 er tilsvarende BenchmarkBSort100 som er kommentert tidligere. 

func benchmarkQsort er tilsvarende benchmarkBSort. Forskjellen er at funksjonen tester ulike funksjoner. 

func BenchmarkBSortModified100,1000 og 10000 er tilsvarende BenchmarkBSort100 som er kommenert tidligere.

 
 





 





 

 


