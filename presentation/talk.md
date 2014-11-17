## Plan prezentacji
- Zagadnienie parametryzowanej złożoności obliczeniowej i problem pokrycia wierzchołkowego
- Opis zaimplementowanych metod redukcji dziedziny
- Przeprowadzone badania

## Parametryzowana złożoność obliczeniowa
- Problemy klasy $\mathcal{FPT}$ - podzbiór problemów klasy $\mathcal{NP}$-trudnych.
- Możliwość zmniejszenia złożoności obliczeniowej do wielomianowej względem rozmiaru danych $n$ oraz wykładniczej lub gorszej względem pewnego parametru $k$.
- Aktualna dziedzina algorytmiki, poszukiwana jest odpowiedź na pytanie czy $\mathcal{P}=\mathcal{NP}$.
- Problem pokrycia wierzchołkowego należy do klasy $\mathcal{FPT}$.

## Problem pokrycia wierzchołkowego
- Problem decyzyjny.
- Odpowiedź na pytanie czy dla pewnego grafu $G=(V, E)$ istnieje zbiór wierzchołków $C \subseteq V$ pokrywajacych łącznie każdą jego krawędź.
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
- Redukcja w oparciu o przepływ w sieci.
- Redukcja w oparciu o usuwanie koron z dziedziny.
- Algorytm Chena, Kanji oraz Xia.

## Redukcja w oparciu o przepływ w sieci
Kroki algorytmu:

- Konstrukcja grafu dwudzielnego i przeformułowanie do sieci przepływowej.
- Znalezienie maksymalnego przepływu.
- Wyznaczenie jądra dziedziny zgodnie z twierdzeniem Nemhausera--Trottera, usunięcie pozostałych wierzchołków.
- $O(n^{5/2})$ dla algorytmu Dinica, $O(n^5)$ dla algorytmu Edmondsa--Karpa.

## Redukcja w oparciu o przepływ w sieci
### Przykład
\includegraphics[width=0.25\textwidth,height=11em]{images/nf-before.pdf}
\includegraphics[width=0.5\textwidth,right=0.75\textwidth]{images/nf-bipartite.pdf}



- Graf dwudzielny stworzony jest według specyficznych reguł.