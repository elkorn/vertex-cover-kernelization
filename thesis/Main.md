\RequirePackage{ifpdf}
\newif\ifelektroniczna
\newif\ifjednostronna

%\elektronicznatrue
\elektronicznafalse

% \jednostronnafalse
\jednostronnatrue

\ifjednostronna
\def\strony{oneside,openany}
\else
\def\strony{twoside,openright}
\fi

\ifpdf
\documentclass[pdftex,12pt,a4paper,\strony,colorlinks,nocenter,noupper,crosshair]{thesis}
\usepackage[pdftex]{graphicx}
\usepackage[pdftex]{hyperref}
\hypersetup{colorlinks,%
  citecolor=black,%
  filecolor=black,%
  linkcolor=black,%
  urlcolor=black,%
pdftex}
\pdfcompresslevel=1
\else
\documentclass[12pt,a4paper,\strony,nocenter,noupper,crosshair]{thesis}
\usepackage{graphicx}
\fi

\usepackage{url}
\usepackage{TitlePage}
\usepackage{theoremref}
\usepackage{mathtools}
\usepackage{tikz}
\usepackage{algorithm}% http://ctan.org/pkg/algorithms
\usepackage{algpseudocode}% http://ctan.org/pkg/algorithmicx
\usepackage{enumerate}
\makeatletter
\renewcommand{\ALG@name}{Pseudokod}
\renewcommand{\listalgorithmname}{Zbiór pseudokodów}
\renewcommand{\algorithmicrequire}{\textbf{Argumenty: }}
\renewcommand{\algorithmicensure}{\textbf{Wynik: }}
\makeatother
\usetikzlibrary{calc, shapes, backgrounds}
\usepackage[utf8]{inputenc}
\def\rodzaj{Praca magisterska}

\def\wydzial{Automatyki, Elektroniki i~Informatyki}

\def\tytul{Parametryzowana złożoność obliczeniowa}
\def\tytulpdf{Parametryzowana złożoność obliczeniowa}

\def\autor{Autor: Korneliusz Adam Caputa}
\def\promotor{Kierujący pracą: prof.~dr~hab.~inż. Zbigniew Czech}
\def\konsultant{}
\def\data{Gliwice, listopad 2014}
\def\slowakluczowe{Fixed point tractability,graph theory,vertex cover,kernelization}

\graphicspath{{./pictures/}}

\DeclareUnicodeCharacter{00A0}{~}

\ifpdf
\ifelektroniczna
\usepackage[
  pdfusetitle=true,
  pdfsubject={\tytulpdf},
  pdfkeywords={\slowakluczowe},
  pdfcreator={\autor},
  pdfstartview=FitV,
  linkcolor=blue,
  citecolor=red,
]{hyperref}
\fi
\fi

\usepackage{layout}

\usepackage{t1enc,amsmath}
\usepackage[OT4,plmath]{polski}
\usepackage{helvet}

%\usepackage{anysize}
%\marginsize{3cm}{2.5cm}{2.5cm}{2.5cm}%LPGD
%\setlength{\textheight}{24cm}
%\usepackage{multirow}
%\ifpdf\usepackage{pdflscape}\else\usepackage{lscape}\fi
%\usepackage{longtable}
%\usepackage{geometry}
%GATHER{thesis.bib}
%\usepackage[twoside]{geometry}
%\geometry{ lmargin=3.5cm, rmargin=2.5cm, tmargin=3cm, bmargin=3cm,
%headheight=1cm, headsep=0.5cm, footskip=0pt }
%\def\fixme#1{}f

\textwidth 150mm
\textheight 225mm
\usepackage{amsfonts}
\usepackage{amsthm}
\usepackage{subfig}
\captionsetup[subfigure]{justification=centerfirst}
\usepackage{cite}
\usepackage{listings}
\lstset{
  language={Go},
  captionpos=b,
  inputencoding=utf8
}
\usepackage{cleardpempty}
\usepackage{float}
\usepackage{textcomp}
\def\vec#1{\ensuremath{\mathbf{#1}}}
\def\ang#1{ang.~\emph{#1}}
\def\lat#1{lac.~\emph{#1}}
\def\e{\ensuremath{\textrm{\normalfont{}e}}}
\def\degree{\ensuremath{^{\circ}}\protect}
\def\fixme#1{\marginpar{\tiny{}#1}}
\def\labelitemi{--}
\def\labelitemii{--}
\def\labelitemiii{--}

\newtheorem{theorem}{\parindent=0pt{\textbf{Twierdzenie}}}[section]
\newtheoremstyle{named}{}{}{\itshape}{}{\bfseries}{.}{.5em}{\thmname{#1}\thmnumber{ #2}\textnormal{\thmnote{#3}}}
\theoremstyle{named}
\newtheorem*{namedtheorem}{Twierdzenie}
\newtheorem{property}{\parindent=0pt{\textbf{Własność}}}[section]
\newtheorem{lemma}{\parindent=0pt{\textbf{Lemat}}}[section]
\newtheorem{corollary}{\parindent=0pt{\textbf{Wniosek}}}[section]
\newtheorem{definition}{\parindent=0pt{\textbf{Definicja}}}[section]
\newtheorem{proposition}{\parindent=0pt{\textbf{Założenie}}}[section]
\newtheorem{conjecture}{\parindent=0pt{\textbf{Przypuszczenie}}}[section]
\newtheorem{note}{\parindent=0pt{\textbf{Uwaga}}}[chapter]
\newenvironment{bproof}{\parindent=0pt{\textbf{Dowód.} }}{\begin{flushright}$\square$\end{flushright}}

  \def\captionlabeldelim{.}
  \linespread{1}
  \chapterfont{\Huge\bfseries}
  \sectionfont{\bfseries\Large}
  \subsectionfont{\bfseries\large}
  \institutionfont{\bfseries}%\mdseries}
  \def\captionlabelfont{\bfseries}

  \renewcommand{\figureshortname}{Rys.}
  \renewcommand{\tableshortname}{Tab.}

  \renewcommand\floatpagefraction{.9}
  \renewcommand\topfraction{.9}
  \renewcommand\bottomfraction{.9}
  \renewcommand\textfraction{.1}
  \setcounter{totalnumber}{50}
  \setcounter{topnumber}{50}
  \setcounter{bottomnumber}{50}

  \newcommand{\topcaption}{%
    \setlength{\abovecaptionskip}{0pt}%
    \setlength{\belowcaptionskip}{10pt}%
  \caption}

  \newcommand\scalemath[2]{\scalebox{#1}{\mbox{\ensuremath{\displaystyle #2}}}}

  \hyphenation{wew-nętrz-nej}

  % \makeatletter
  % \renewcommand\fs@ruled{\def\@fs@cfont{\bfseries}\let\@fs@capt\floatc@ruled
  %   \def\@fs@pre{\hrule height1.0pt depth0pt \kern2pt}
  %   \def\@fs@post{\vskip-1.5\baselineskip\kern2pt\hrule\relax}%
  %   \def\@fs@mid{\kern2pt\hrule\kern2pt}
  % \let\@fs@iftopcapt\iftrue}
  % \makeatother

  \floatstyle{ruled}
  % \newfloat{sample}{thp}{lop}
  % \floatname{sample}{Przykład}

  \begin{document}

  \bibliographystyle{unsrt}
  \frontmatter
  \stronatytulowa
  %\cleardoublepage
  %\maketitle
  %\tocbibname

  \tableofcontents % \listoffigures \listoftables \listof{sample}{Spis przykładów}
  %\listofacros
  %\input{abbrev_body}
  %\newpage
  %\input{spis_oznaczen}

  \mainmatter
  \input{Introduction}
  \input{Theory}
  \input{Internals}
  \input{Results}
  \input{Summary}

  \addcontentsline{toc}{chapter}{\bibname}
  % \input{Appendixes}
  \bibliography{Main}

  % \renewcommand{\appendixname}{Dodatek}
  % \appendix

  \end{document}
