mlr --from regtest/input/ten.dkvp head -n 1 then put -q 'end { @v=[1,2,3,4,5]; unset @v[2]; dump @v }'
