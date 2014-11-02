setwd("/home/elkorn/Code/go/src/github.com/elkorn/vertex-cover-kernelization/main/measurement-results")

naive <- read.table("MeasureNaive", header=TRUE, sep="\t")
cr <- read.table("results_vc_cr", header=TRUE, sep="\t")
nf <- read.table("results_vc_nf", header=TRUE, sep="\t")
bnb <- read.table("results_vc_bnb", header=TRUE, sep="\t")
maxnum <- function(col) {
  return(max(sapply(col, as.numeric)))
}

lwd = 2

yMax = max(maxnum(bnb$Ts), maxnum(nf$Ts), maxnum(cr$Ts))
xMax = max(length(bnb$Ts), length(nf$Ts), length(cr$Ts))
x <- seq(100, 2000, 100)
c <- rainbow(4)

plot(x, x, xlab = "Liczba wierzchołków grafu", ylab="Czas wyznaczania pokrycia wierzchołkowego [s]", main="Wyznaczanie pokrycia wierzchołkowego z przetwarzaniem wstępnym", ylim=c(0, yMax), pch = 20, type="n")
mkLineX <- function(x, data, idx) {
  lines(x, data$Ts, col=c[idx], lwd = lwd, type="o")
}

mkLine <- function(data, idx) {
  n = length(data$Ts)
  print(n)
  mkLineX(x[1:n], data, idx)
}


mkLine(bnb, 1)
mkLine(cr, 2)
mkLine(nf, 3)
legend(100,yMax, c("Brak kernelizacji","Redukcja koron", "Przepływ w sieci", "Metoda siłowa"), lty = c(1,1,1,1), lwd=c(lwd,lwd,lwd,lwd),col=c)       

mkLineX(naive$V[1:length(naive$Ts)], naive, 4)