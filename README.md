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
I denne filen ligger det en func som heter main. Denne funksjonen gjør en print av teksten i funksjonen. 

Sum.go
Denne filen inneholder fem funksjoner. Alle returnerer a+b, som er summen av tallene som blir skrevet inn i terminalen. 
Det er fem ulike funksjoner fordi programmet skal kunne opperere med forskjellige typer, og de er som følger:int8,
float64, uint32, int32 og int64. 
Int8 var ikke med i oppgaven, men vi tok dette med fordi vi ville bruke den i testen for å få bedre forståelse for testene, 
samt vi ville kunne forstå rangen bedre med lave tall før vi gikk videre til de større tallene. 

Main_sum.go
Her finner vi mainen til sum.go. Denne har en func som heter main, som kjører programmet. 
I starten av funksjonen har vi variablen "finish", som blir brukt som en boolean i en for-loop. Denne loopen holder programmet kjørende,
helt til den brytes, og programmet avsluttes. Etter det blir det deklarert en variabel, typeVal, som er en string. Den bestemmer hvilken 
type som skal kalkuleres. Så kommer det en to prints, som sier hvilken typer som er valide, og at man skal skrive hvilken type man vil kalkulere.
Det man skriver i terminalen blir så scannet, og den leter etter typeVal.
Videre kommer det fire if statements, hvor koden er lik, forskjellen er bare hvilken type det dreier seg om, for eksemple uint32 eller int32. 
De tar den riktige typen, utifra hva som ble skrevet i terminalen,scanner tallene som blir skrevet inn, så blir metoden som hører til typen kalt 
på fra sum.go, som legger sammen de to tallene. Det blir brukt to variabler, val og val2. 
Når programmet har lagt sammen to tall kommer det en print som forklarer hvordan man avslutter programmet. Hvis man skriver "yes", blir finnish==true og
for loopen brytes, og programmet avsluttes. Hvis man skriver "no" kan man fortsette med å legge sammen tall. 

sum_test.go
Et testprogram som tester at kalkulatoren "sum" fungerer som den skal.
Programmet tester med de fem ulike typene. En funksjon for hver type, med en tilhørende struct. 
I structen er det tre variabler, to for tallene som blir lagt sammen, og én for tallet som er forventet å få som svar. 
Tallene som blir testet står også i structen. 
I funksjonene er det en for loop med variablen "v", som bestemmer hvilken range som blir brukt, og rangen er lik som tallene i structen. 
Så kommer en if statement som sier at hvis for eksempel if val:= SumInt8 ikke er lik det forventede tallet, printer programmet en error statement, 
som sier hva som ble kalkulert, og hva som var forventet. 

boring.go
Programmet inneholder to funksjoner, Boring01 og Boring10.
Boring01 har "msg" som parameter. Boring01 har en for loop som ikke har noen break. Variablen i begynner på null, og så teller den oppover. 
Funksjonen printen en melding eller "msg", og variablen i som stiger med ett nummer for hver linje som blir printet. 
Time.sleep gjør at den bruker litt tid på å printe ut neste linje. 
Den andre funcen, Boring10, som har "msg" og "c" som parametere. 
Denne funcen har en lik for loop som Boring01.
så kommer det en fmt.sprintf som konverterer meldingen til en string fra hvilken som helst type.
Time.sleep gjør at den bruker litt tid på å printe ut neste linje. 
C Channel venter på svar fra sprintf. 



Boring_main.go
Dette er den ene mainen til boring.go. 
Dette programmet kjører boring.go, og funksjonen boring01 blir kalt på fra boring.go. 

main_boring_goroutine.go
Dette er den andre mainen til boring.go. 
Funksjonen boring10 blir kalt på fra boring.go. 
En channel blir laget og har variabelen c. 
Det er en for loop som som teller oppover fra 0, som er variablen "i".
En melding blir printet.
Boring10 kombinerer teksten som vi sender inn med integeren i loopen på boring10.  


sorting.go
Et program som sorterer integers i stigende rekkefølge. 
func bubble_sort_modified har en liste av integers som parameter. 
Variabelen "n" er lik lengde på listen. 
Variabelen swapped == true, en boolean som er sann når integersene er sortert i riktig rekkefølge. 
En for-loop som sier at når swapped==false skal den iterere gjennom tallene, variablen "i" = 1.
Når "i" er mindre enn lengden på listen minus 1, skal loopen fortsette. 
If listen [i-1] er større enn listen [i], blir verdiene i tallrekken blir sortert. 
Hvis list[i-1] = list[i-1] er swapped true, og for loopen brytes. 

func bubble_sort har også en liste av integers som parameter. 
Variabelen n = lengden på listen. 
En for-loop med variabelen i:=0, når "i" er mindre enn "n"(lengden på listen), fortsetter "i" videre på neste tall. 
En for-loop med variabelen j:=0, når "j" er mindre enn "n"(lengden på lisetn) minus 1, fortsetter "j" videre på neste tall. 
if list[j] er større enn list [j+1] blir temp:= list[j+1]. 
list [j+1] er da lik list[j]. 
Og det betyr at list[j] er lik temp. 

func Qsort at values som er integers i parameter. 
Denne funksjonen kaller på qsort og har parameterene values, l:0 og lendgden på listen som er values -1.

func qsort har tre parametere, values av typen int, l av typen int, og r av typen int. 
if "l" er større eller lik "r", returnerer den en blank return hvis if statementen stemmer. 
Variablen pivot:= values[l]. 
variabelen i:= "l" pluss 1. 
En for-loop med variabelen j:=l. Hvis "j" er mindre eller lik "r" går j videre på neste tall. 
if pivot er større enn values[j], og values[i], values[j] er lik values [j], values [i], går i videre på neste tall. 
Den kjører en loop til loopen er komplett, og når den er komplett returnerer den blankt. 

sorting_test.go
Et program som kjører benchmark tester og test på funksjoner fra sorting.go.  

 
func init setter den tiden til systemet til seed, og den sikrer at resultatet blir likt hver gang. 

func perm gjør at en får tilfeldige tall å teste med. 

func BenchmarkBsort100 har b*testing.B som parameter, som gjør det mulig å kjøre benchmark tester. 
Den kaller på benchmarkBsort funcen som har en int og b som parameter. integeren "i", er hvor mange tall som er listen
som skal sorteres, i dette tilfellet 100. 
Denne funcen gjentar seg flere ganger, bare at variabelen "i" endres til 1000, og 10 000. 

func benchmarkBSort har parameterene "i" og "b". "i" er hvor mange tall som er listen som skal sorteres, og "b" kommer
fra testing packagen som er importert. 
Det er en for-loop med variabelen j:=0, når j er mindre enn b.N fortsetter j til neste tall. 
b.StopTimer gjør at testen blir stoppet.
b.startTimer starter testen igjen. 
Dette er på grunn av tiden det tar å hente verdiene som skal testes ikke skal bli tatt med i testn

func benchmarkBSortModified er lik som benchmarkBSort, bare at den er for to ulike funksjoner i sorting.go.  

func TestQSort tester QSort funksjonen fra sorting.go. 
Den har variablene values og expected, og begge holder en liste med tall. 
values er ikke sortert, mens expected er tallene sortert i stigende rekkefølge. 
Hvis QSort ikke har sortert de i riktig rekkefølge blir det printet en feilmelding. 

func BenchmarkQSort100 er tilsvarende BenchmarkBSort100 som er kommentert tidligere. 

func benchmarkQsort er tilsvarende benchmarkBSort. Forskjellen er at funksjonen tester ulike funksjoner. 

func BenchmarkBSortModified100,1000 og 10000 er tilsvarende BenchmarkBSort100 som er kommenert tidligere.

 
 





 





 

 


