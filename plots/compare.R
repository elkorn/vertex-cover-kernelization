setwd("/home/elkorn/Code/go/src/github.com/elkorn/vertex-cover-kernelization/main")

bnb <- read.table("results_vc_bnb", header=TRUE, sep="\t")
cr <- read.table("results_vc_cr", header=TRUE, sep="\t")
nf <- read.table("results_vc_nf", header=TRUE, sep="\t")

maxnum <- function(col) {
  return(max(sapply(col, as.numeric)))
}

lwd = 2

yMax = max(maxnum(bnb$Ts), maxnum(nf$Ts), maxnum(cr$Ts))
x <- bnb$V#range(c(1, 1600), 10)
c <- rainbow(3)
# lm.out = lm(y ~ x)
# plot(y ~ x)
# abline(lm.out, col="red")
# set up the plot 
# plot(x, y, xlab = "Liczba wierzchołków grafu", ylab="Czas wyznaczania pokrycia wierzchołkowego", pch = 20, type="n") 
# lines(x, y, col="red", lwd = 1.5, type="o")
# symbols(x=sorted$V, y=sorted$E, circles=crown$T, inches=1/3, ann=F, bg="steelblue2", fg=NULL)

plot(x, bnb$Ts, xlab = "Liczba wierzchołków grafu", ylab="Czas wyznaczania pokrycia wierzchołkowego [s]",ylim=c(0, yMax), pch = 20, type="n")
mkLine <- function(data, idx) {
  lines(x[1:16], data$Ts[1:16], col=c[idx], lwd = lwd, type="o")
}


mkLine(bnb, 1)
mkLine(cr, 2)
mkLine(nf, 3)
legend(100,yMax, c("Brak kernelizacji","Redukcja koron", "Przepływ w sieci"), lty = c(1,1,1), lwd=c(lwd,lwd,lwd),col=c)       

