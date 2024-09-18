 
 # Build check

 kustomize build . 
 
  # Build check and apply
 kustomize build . | kubectl apply -f -

 or

 kubectl apply -k . 