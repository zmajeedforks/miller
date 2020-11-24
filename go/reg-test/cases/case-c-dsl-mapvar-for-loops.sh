run_mlr --from $indir/abixy put -q '
  func f() {
    return {
      "a:".$a: $i,
      "b:".$b: $y,
      NR: NF,
    }
  }
  for (k,v in f()) {
    print "k=".k;
    print "v=".v;
  }
  print
'
