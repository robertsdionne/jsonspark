# jsonspark

Command line tool for depicting numerical values within sequences of JSON data as spark lines in the terminal.

The following JSON:
```JSON
[
  {
    "averageLoss": 1.3072037222204602,
    "accuracy": 0.7961115007048624
  },
  {
    "averageLoss": 1.0950766502916955,
    "accuracy": 0.8284353658266508
  },
  {
    "averageLoss": 1.0659628237911414,
    "accuracy": 0.8341821064915779
  },
  {
    "averageLoss": 1.0561794574499441,
    "accuracy": 0.8361791368768862
  },
  {
    "averageLoss": 1.0519946308009303,
    "accuracy": 0.8371011584581296
  },
  {
    "averageLoss": 1.0507322509047852,
    "accuracy": 0.8374584676416922
  },
  {
    "averageLoss": 1.0509473261079354,
    "accuracy": 0.8376648443177903
  },
  {
    "averageLoss": 1.0513566022246583,
    "accuracy": 0.8379441199332248
  },
  {
    "averageLoss": 1.051017922874153,
    "accuracy": 0.8381535767726236
  },
  {
    "averageLoss": 1.0509318358121373,
    "accuracy": 0.8383845952442855
  },
  {
    "averageLoss": 1.050666877902315,
    "accuracy": 0.8385550356174658
  },
  {
    "averageLoss": 1.0504042720981572,
    "accuracy": 0.8387182886937977
  },
  {
    "averageLoss": 1.0511238890230516,
    "accuracy": 0.8384277192320015
  },
  {
    "averageLoss": 1.0514740070664277,
    "accuracy": 0.8385273134709509
  }
]
```

will generate this JSON summary:
```JSON
{
  "accuracy": "▁▇████████████",
  "averageLoss": "█▂▁▁▁▁▁▁▁▁▁▁▁▁"
}
```
