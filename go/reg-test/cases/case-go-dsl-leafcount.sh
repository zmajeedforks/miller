run_mlr --opprint --from $indir/s.dkvp put '
  $lcx = leafcount($x);
  $lcn = leafcount($nonesuch);
  $lca1 = leafcount([1,2,3]);
  $lca2 = leafcount([1,[4,5,6],3]);
  $lca3 = leafcount([1,{"s":4,"t":[7,8,9],"u":6},3]);
  $lcm1 = leafcount({"s":1,"t":2,"u":3});
  $lcm2 = leafcount({"s":1,"t":[4,5,6],"u":3});
  $lcm3 = leafcount({"s":1,"t":[4,{"x":8, "y": 9},6],"u":3});
'
