# Getting started with your own model

This guide will show you how package up your trained model with Cog and run predictions on it.

## Define the environment

First step is to create a `cog.yaml`. It defines all the different things that need to be installed for your model to run. You can think of it as a simple way of defining a Docker image.

For example:

```yaml
environment:
  python_version: "3.8"
  python_packages:
    - "torch==1.7.0"
```

This will generate a Docker image with Python 3.8 and PyTorch 1.7 installed, for both CPU and GPU, with the correct version of CUDA, and various other sensible best-practices.

To run a command inside this environment, prefix it with `cog run`:

```
$ cog run python
✓ Building Docker image from environment in cog.yaml... Successfully built 8f54020c8981
Running 'python' in Docker with the current directory mounted as a volume...
────────────────────────────────────────────────────────────────────────────────────────

Python 3.8.10 (default, May 12 2021, 23:32:14)
[GCC 9.3.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

This is handy for ensuring a consistent environment for development or training.

With `cog.yaml`, you can also install system packages and other things. [Take a look at the full reference to see what else you can do.](yaml.md)

## Define how to run predictions

Next step is to the define the interface for running predictions on your model. It looks something like this:

```python
import cog
from pathlib import Path
import torch

class Model(cog.Model):
    # Load the model into memory to make running multiple inferences efficient
    def setup(self):
        self.net = torch.load("weights.pth")

    # Define the arguments and types the model takes as input
    @cog.input("input", type=Path, help="Image to enlarge")
    @cog.input("scale", type=float, default=1.5, help="Factor to scale image by")
    # Run a single prediction on the model
    def predict(self, input, scale):
        # ... pre-processing ...
        output = self.net(input)
        # ... post-processing ...
        return output
```

Put this in a file called `predict.py` and fill in the functions with your own model's setup and prediction code. You might want to import parts of your model from another file.

You also need to define the inputs to your model using the `@cog.input()` decorator, as demonstrated above. The first argument maps to the name of the argument in the `predict()` function, and it also takes these other arguments:

- `type`: Either `str`, `int`, `float`, `bool`, or `Path` (be sure to add the import, as in the example above). `Path` is used for files. For more complex inputs, save it to a file and use `Path`.
- `help`: A description of what to pass to this input for users of the model
- `default`: A default value to set the input to. If this argument is not passed, the input is required. If it is explicitly set to `None`, the input is optional.
- `min`: A minimum value for `int` or `float` types.
- `max`: A maximum value for `int` or `float` types.
- `options`: A list of values to limit the input to. It can be used with `str`, `int`, and `float` inputs.

For more details about writing your model interface, [take a look at the prediction interface documentation](python.md).

Next, add this line at the top of your `cog.yaml` file so Cog knows how to run predictions:

```yaml
model: "predict.py:Model"
```

That's it! To test this works, try running a prediction on the model:

```
$ cog predict -i @input.jpg
✓ Building Docker image from environment in cog.yaml... Successfully built 664ef88bc1f4
✓ Model running in Docker image 664ef88bc1f4

Written output to output.png
```

## Next steps

You might want to look at [the main getting started guide](getting-started.md), where it shows you how to set up a Cog server and push your model to it.
