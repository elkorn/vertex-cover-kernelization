## Plan prezentacji
- Zagadnienie parametryzowanej złożoności obliczeniowej i problem pokrycia wierzchołkowego.
- Opis zaimplementowanych metod przetwarzania wstępnego i redukcji dziedziny.
- Opis algorytmu Chena, Kanji oraz Xia.
- Omówienie przeprowadzonych badań.
- Podsumowanie i wnioski.

## Parametryzowana złożoność obliczeniowa
- Problemy klasy $\mathcal{FPT}$ stanowią podzbiór problemów klasy $\mathcal{NP}$-trudnych.
- Możliwość zmniejszenia złożoności obliczeniowej do wielomianowej względem rozmiaru danych $n$ oraz wykładniczej lub gorszej względem pewnego parametru $k$.
- Aktualna dziedzina algorytmiki, poszukiwana jest odpowiedź na pytanie czy $\mathcal{P}=\mathcal{NP}$.
- Problem pokrycia wierzchołkowego należy do klasy $\mathcal{FPT}$.

## Problem pokrycia wierzchołkowego
- Problem decyzyjny.
- Odpowiedź na pytanie czy dla pewnego grafu $G=(V, E)$ istnieje zbiór wierzchołków $C \subseteq V$ pokrywających łącznie każdą jego krawędź.
- Parametryzacja: poszukiwane jest pokrycie o rozmiarze $|C| \leq k$.

## Problem pokrycia wierzchołkowego
### Przykład: graf Petersena
\includegraphics[width=0.5\textwidth,natwidth=692,natheight=100]{images/petersen.pdf}
\includegraphics[width=0.5\textwidth,natwidth=692,natheight=100]{images/petersen-cover.pdf}

- Istnieje pokrycie wierzchołkowe $C$ dla parametru $k=6$.

## Zaimplementowane algorytmy
### Przetwarzanie wstępne
Proste algorytmy redukujące dziedzinę poszukiwań --- $O(n^2)$.

- Usunięcie wierzchołków izolowanych.
- Usunięcie wierzchołków stopnia 1.
- Usunięcie wierzchołków stopnia 2 o rozłącznym sąsiedztwie.
- Zwinięcie wierzchołków stopnia 2 wraz z połączonym sąsiedztwem.

### Redukcja dziedziny
- Usunięcie z dziedziny wierzchołków stopnia $d(v) \geq k$.
- Redukcja dziedziny w oparciu o przepływ w sieci.
- Redukcja dziedziny w oparciu o usuwanie koron.

### Algorytm Chena, Kanji oraz Xia

## Redukcja dziedziny w oparciu o przepływ w sieci
Kroki algorytmu:

- Konstrukcja grafu dwudzielnego i przeformułowanie do sieci przepływowej.
- Znalezienie maksymalnego przepływu.
- Wyznaczenie jądra dziedziny zgodnie z twierdzeniem Nemhausera--Trottera, usunięcie pozostałych wierzchołków.
- $O(n^{5/2})$ dla algorytmu Dinica, $O(n^5)$ dla algorytmu Edmondsa--Karpa.

## Redukcja w oparciu o przepływ w sieci
### Przykład
\includegraphics[width=0.25\textwidth,height=11em]{images/nf-before.pdf}
\includegraphics[width=0.5\textwidth,right=0.75\textwidth]{images/nf-bipartite2.pdf}

- Graf dwudzielny stworzony jest według specyficznych reguł.

## Redukcja dziedziny przez usunięcie koron
- Korona to para zbiorów wierzchołków $(I, H)$.
    - $I$ stanowi zbiór niezależny.
    - $H$ (głowa korony) to sąsiedztwo wierzchołków zbioru $I$.
    - Istnieje skojarzenie $M$ na krawędziach pomiędzy zbiorami $I$ oraz $H$, które zawiera zbiór $H$.
    - Zachodzi $|H| \leq |I|$.

## Redukcja dziedziny przez usunięcie koron
### Przykład
\includegraphics[width=0.5\textwidth]{images/crown-before.pdf}
\includegraphics[width=0.5\textwidth]{images/crown.pdf}

- Zbiór $I$ zawiera wierzchołki nienależące do maksymalnego skojarzenia grafu.
- Złożoność zależy od algorytmu wyznaczającego maksymalne skojarzenie:
    - $O(n^{5/2})$ dla algorytmu Dinica.
    - $O(n^{4})$ dla algorytmu skurczania kwiatów Edmondsa.

## Algorytm Chena, Kanji oraz Xia

- Pełne rozwiązanie problemu pokrycia wierzchołkowego.
- Bardzo skomplikowany analitycznie.
- Rekurencyjny algorytm z rozgałęzieniami.
- Każda iteracja wpierw redukuje graf przez redukcję koron oraz prawie-koron, a następnie identyfikuje i usuwa struktury ,,pożyteczne'' z dziedziny.
    - Istnieje 12 rodzajów tychże struktur.
    - Kolejne iteracje operują na zawężonych dziedzinach.

- $O(1.2738^k + kn)$ --- obecnie jeden z najszybszych na świecie.

## Badania eksperymentalne
### Dane testowe
- Wygenerowano zbiór grafów losowych o 100--2000 wierzchołkach z krokiem stu wierzchołków.
- Stopień $d(v)$ każdego z wierzchołków wybierany był losowo z przedziału $<1, 5>$ zgodnie z rozkładem prawdopodobieństwa $P(d(v)) \propto \frac{1}{d(v)}$.

## Badania eksperymentalne
### Wykonane pomiary
- Wyniki uzyskano przez zmierzenie czasu działania następujących operacji:
    1. Rozwiązanie problemu pokrycia wierzchołkowego metodą siłową.
    1. Rozwiązanie problemu pokrycia wierzchołkowego metodą drzewa poszukiwań z ograniczeniami:
        1. Na dziedzinie pierwotnej.
        1. Na dziedzinie zawężonej algorytmem opartym o przepływ w sieci.
        1. Na dziedzinie zawężonej algorytmem redukcji koron.
    1. Przetwarzanie wstępne grafu.
- Pomiary z punktu 2 zrealizowano zarówno dla dziedziny pierwotnej jak i po uprzednim przetwarzaniu wstępnym.

## Podsumowanie i wnioski

- Zaimplementowano pięć algorytmów związanych z rozwiązywaniem problemu pokrycia wierzchołkowego.
- Wyniki badań pokazują pozytywny wpływ algorytmów redukcji dziedziny oraz przetwarzania wstępnego na czas rozwiązywania problemu.
- Stworzone rozwiązania wymagają optymalizacji w celu umożliwienia zastosowań praktycznych.
- Osiągnięto założony cel drugorzędny --- stworzono elastyczną implementację struktury grafu ułatwiającą prowadzenie dalszych badań nad innymi metodami rozwiązywania problemu pokrycia wierzchołkowego w czasie wielomianowym.

