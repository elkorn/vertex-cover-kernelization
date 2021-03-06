setwd("/home/elkorn/Code/go/src/github.com/elkorn/vertex-cover-kernelization/main/measurement-results")

# naive <- read.table("MeasureNaive", header=TRUE, sep="\t")
cr <- read.table("results_vc_cr", header=TRUE, sep="\t")
nf <- read.table("results_vc_nf", header=TRUE, sep="\t")
bnb <- read.table("results_vc_bnb", header=TRUE, sep="\t")
maxnum <- function(col) {
  return(max(sapply(col, as.numeric)))
}

lwd = 2

yMax = max(maxnum(bnb$Ts), maxnum(nf$Ts), maxnum(cr$Ts))
x <- seq(100, 2000, 100)
c <- rainbow(4)

plot(x, seq(0, yMax, 10), xlab = "Liczba wierzchołków grafu", ylab="Czas wyznaczania pokrycia wierzchołkowego [s]", main="Wyznaczanie pokrycia wierzchołkowego z przetwarzaniem wstępnym", ylim=c(0, yMax), pch = 20, type="n")
mkLine <- function(data, idx) {
  n = length(data$Ts)
  print(n)
  lines(x[1:n], data$Ts, col=c[idx], lwd = lwd, type="o")
}


mkLine(bnb, 1)
mkLine(cr, 2)
mkLine(nf, 3)
legend(100,yMax, c("Brak kernelizacji","Redukcja koron", "Przepływ w sieci", "Metoda siłowa"), lty = c(1,1,1,1), lwd=c(lwd,lwd,lwd,lwd),col=c)       

lines(seq(10, 30, 10), seq(100, 1000, 100), col=c[3], lwd = lwd, type="o")
