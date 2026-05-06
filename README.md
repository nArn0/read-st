read-st
=======

No Python, no problem. This binary just take a safetensors file as argument and list all tensors present in the file.

On windows, you can just drag and drop the safetensors file on the binary to get the list.

You can also send the result to a text file using this:

```
read-st model.safetensors > layers.txt
```

PS: the binary will wait for a 'Enter' keypress at the end to allow reading the result if the drag and drop method is used.
