// init
@256 // 0
D=A // 1
@SP // 2
M=D // 3
// call Sys.init 0
@return.1 // 4
D=A // 5
@0 // 6
A=M // 7
M=D // 8
@0 // 9
M=M+1 // 10
@LCL // 11
D=M // 12
@0 // 13
A=M // 14
M=D // 15
@0 // 16
M=M+1 // 17
@ARG // 18
D=M // 19
@0 // 20
A=M // 21
M=D // 22
@0 // 23
M=M+1 // 24
@THIS // 25
D=M // 26
@0 // 27
A=M // 28
M=D // 29
@0 // 30
M=M+1 // 31
@THAT // 32
D=M // 33
@0 // 34
A=M // 35
M=D // 36
@0 // 37
M=M+1 // 38
@0 // 39
D=M // 40
@0 // 41
D=D-A // 42
@5 // 43
D=D-A // 44
@ARG // 45
M=D // 46
@0 // 47
D=M // 48
@LCL // 49
M=D // 50
@Sys.init // 51
0;JMP // 52
(return.1)
// function Class1.set 0
(Class1.set)
// C_PUSH argument:0
@ARG // 53
D=M // 54
@0 // 55
D=D+A // 56
A=D // 57
D=M // 58
@SP // 59
A=M // 60
M=D // 61
@0 // 62
M=M+1 // 63
// C_POP static:0
@Class1.vm.0 // 64
D=A // 65
@0 // 66
D=D+A // 67
@R13 // 68
M=D // 69
@0 // 70
M=M-1 // 71
A=M // 72
D=M // 73
@R13 // 74
A=M // 75
M=D // 76
// C_PUSH argument:1
@ARG // 77
D=M // 78
@1 // 79
D=D+A // 80
A=D // 81
D=M // 82
@SP // 83
A=M // 84
M=D // 85
@0 // 86
M=M+1 // 87
// C_POP static:1
@Class1.vm.1 // 88
D=A // 89
@0 // 90
D=D+A // 91
@R13 // 92
M=D // 93
@0 // 94
M=M-1 // 95
A=M // 96
D=M // 97
@R13 // 98
A=M // 99
M=D // 100
// C_PUSH constant:0
@0 // 101
D=A // 102
@0 // 103
A=M // 104
M=D // 105
@0 // 106
M=M+1 // 107
// return
@LCL // 108
D=M // 109
@R13 // 110
M=D // 111
@R13 // 112
D=M // 113
@5 // 114
D=D-A // 115
A=D // 116
D=M // 117
@R14 // 118
M=D // 119
@0 // 120
M=M-1 // 121
A=M // 122
D=M // 123
@ARG // 124
A=M // 125
M=D // 126
@ARG // 127
D=M // 128
@1 // 129
D=D+A // 130
@SP // 131
M=D // 132
@R13 // 133
D=M // 134
@1 // 135
D=D-A // 136
A=D // 137
D=M // 138
@THAT // 139
M=D // 140
@R13 // 141
D=M // 142
@2 // 143
D=D-A // 144
A=D // 145
D=M // 146
@THIS // 147
M=D // 148
@R13 // 149
D=M // 150
@3 // 151
D=D-A // 152
A=D // 153
D=M // 154
@ARG // 155
M=D // 156
@R13 // 157
D=M // 158
@4 // 159
D=D-A // 160
A=D // 161
D=M // 162
@LCL // 163
M=D // 164
@R14 // 165
A=M // 166
0;JMP // 167
// function Class1.get 0
(Class1.get)
// C_PUSH static:0
@Class1.vm.0 // 168
D=A // 169
@0 // 170
D=D+A // 171
A=D // 172
D=M // 173
@SP // 174
A=M // 175
M=D // 176
@0 // 177
M=M+1 // 178
// C_PUSH static:1
@Class1.vm.1 // 179
D=A // 180
@0 // 181
D=D+A // 182
A=D // 183
D=M // 184
@SP // 185
A=M // 186
M=D // 187
@0 // 188
M=M+1 // 189
// sub
@0 // 190
M=M-1 // 191
A=M // 192
D=M // 193
@0 // 194
M=M-1 // 195
A=M // 196
M=M-D // 197
@0 // 198
M=M+1 // 199
// return
@LCL // 200
D=M // 201
@R13 // 202
M=D // 203
@R13 // 204
D=M // 205
@5 // 206
D=D-A // 207
A=D // 208
D=M // 209
@R14 // 210
M=D // 211
@0 // 212
M=M-1 // 213
A=M // 214
D=M // 215
@ARG // 216
A=M // 217
M=D // 218
@ARG // 219
D=M // 220
@1 // 221
D=D+A // 222
@SP // 223
M=D // 224
@R13 // 225
D=M // 226
@1 // 227
D=D-A // 228
A=D // 229
D=M // 230
@THAT // 231
M=D // 232
@R13 // 233
D=M // 234
@2 // 235
D=D-A // 236
A=D // 237
D=M // 238
@THIS // 239
M=D // 240
@R13 // 241
D=M // 242
@3 // 243
D=D-A // 244
A=D // 245
D=M // 246
@ARG // 247
M=D // 248
@R13 // 249
D=M // 250
@4 // 251
D=D-A // 252
A=D // 253
D=M // 254
@LCL // 255
M=D // 256
@R14 // 257
A=M // 258
0;JMP // 259
// function Class2.set 0
(Class2.set)
// C_PUSH argument:0
@ARG // 260
D=M // 261
@0 // 262
D=D+A // 263
A=D // 264
D=M // 265
@SP // 266
A=M // 267
M=D // 268
@0 // 269
M=M+1 // 270
// C_POP static:0
@Class2.vm.0 // 271
D=A // 272
@0 // 273
D=D+A // 274
@R13 // 275
M=D // 276
@0 // 277
M=M-1 // 278
A=M // 279
D=M // 280
@R13 // 281
A=M // 282
M=D // 283
// C_PUSH argument:1
@ARG // 284
D=M // 285
@1 // 286
D=D+A // 287
A=D // 288
D=M // 289
@SP // 290
A=M // 291
M=D // 292
@0 // 293
M=M+1 // 294
// C_POP static:1
@Class2.vm.1 // 295
D=A // 296
@0 // 297
D=D+A // 298
@R13 // 299
M=D // 300
@0 // 301
M=M-1 // 302
A=M // 303
D=M // 304
@R13 // 305
A=M // 306
M=D // 307
// C_PUSH constant:0
@0 // 308
D=A // 309
@0 // 310
A=M // 311
M=D // 312
@0 // 313
M=M+1 // 314
// return
@LCL // 315
D=M // 316
@R13 // 317
M=D // 318
@R13 // 319
D=M // 320
@5 // 321
D=D-A // 322
A=D // 323
D=M // 324
@R14 // 325
M=D // 326
@0 // 327
M=M-1 // 328
A=M // 329
D=M // 330
@ARG // 331
A=M // 332
M=D // 333
@ARG // 334
D=M // 335
@1 // 336
D=D+A // 337
@SP // 338
M=D // 339
@R13 // 340
D=M // 341
@1 // 342
D=D-A // 343
A=D // 344
D=M // 345
@THAT // 346
M=D // 347
@R13 // 348
D=M // 349
@2 // 350
D=D-A // 351
A=D // 352
D=M // 353
@THIS // 354
M=D // 355
@R13 // 356
D=M // 357
@3 // 358
D=D-A // 359
A=D // 360
D=M // 361
@ARG // 362
M=D // 363
@R13 // 364
D=M // 365
@4 // 366
D=D-A // 367
A=D // 368
D=M // 369
@LCL // 370
M=D // 371
@R14 // 372
A=M // 373
0;JMP // 374
// function Class2.get 0
(Class2.get)
// C_PUSH static:0
@Class2.vm.0 // 375
D=A // 376
@0 // 377
D=D+A // 378
A=D // 379
D=M // 380
@SP // 381
A=M // 382
M=D // 383
@0 // 384
M=M+1 // 385
// C_PUSH static:1
@Class2.vm.1 // 386
D=A // 387
@0 // 388
D=D+A // 389
A=D // 390
D=M // 391
@SP // 392
A=M // 393
M=D // 394
@0 // 395
M=M+1 // 396
// sub
@0 // 397
M=M-1 // 398
A=M // 399
D=M // 400
@0 // 401
M=M-1 // 402
A=M // 403
M=M-D // 404
@0 // 405
M=M+1 // 406
// return
@LCL // 407
D=M // 408
@R13 // 409
M=D // 410
@R13 // 411
D=M // 412
@5 // 413
D=D-A // 414
A=D // 415
D=M // 416
@R14 // 417
M=D // 418
@0 // 419
M=M-1 // 420
A=M // 421
D=M // 422
@ARG // 423
A=M // 424
M=D // 425
@ARG // 426
D=M // 427
@1 // 428
D=D+A // 429
@SP // 430
M=D // 431
@R13 // 432
D=M // 433
@1 // 434
D=D-A // 435
A=D // 436
D=M // 437
@THAT // 438
M=D // 439
@R13 // 440
D=M // 441
@2 // 442
D=D-A // 443
A=D // 444
D=M // 445
@THIS // 446
M=D // 447
@R13 // 448
D=M // 449
@3 // 450
D=D-A // 451
A=D // 452
D=M // 453
@ARG // 454
M=D // 455
@R13 // 456
D=M // 457
@4 // 458
D=D-A // 459
A=D // 460
D=M // 461
@LCL // 462
M=D // 463
@R14 // 464
A=M // 465
0;JMP // 466
// function Sys.init 0
(Sys.init)
// C_PUSH constant:6
@6 // 467
D=A // 468
@0 // 469
A=M // 470
M=D // 471
@0 // 472
M=M+1 // 473
// C_PUSH constant:8
@8 // 474
D=A // 475
@0 // 476
A=M // 477
M=D // 478
@0 // 479
M=M+1 // 480
// call Class1.set 2
@return.2 // 481
D=A // 482
@0 // 483
A=M // 484
M=D // 485
@0 // 486
M=M+1 // 487
@LCL // 488
D=M // 489
@0 // 490
A=M // 491
M=D // 492
@0 // 493
M=M+1 // 494
@ARG // 495
D=M // 496
@0 // 497
A=M // 498
M=D // 499
@0 // 500
M=M+1 // 501
@THIS // 502
D=M // 503
@0 // 504
A=M // 505
M=D // 506
@0 // 507
M=M+1 // 508
@THAT // 509
D=M // 510
@0 // 511
A=M // 512
M=D // 513
@0 // 514
M=M+1 // 515
@0 // 516
D=M // 517
@2 // 518
D=D-A // 519
@5 // 520
D=D-A // 521
@ARG // 522
M=D // 523
@0 // 524
D=M // 525
@LCL // 526
M=D // 527
@Class1.set // 528
0;JMP // 529
(return.2)
// C_POP temp:0
@R5 // 530
D=A // 531
@0 // 532
D=D+A // 533
@R13 // 534
M=D // 535
@0 // 536
M=M-1 // 537
A=M // 538
D=M // 539
@R13 // 540
A=M // 541
M=D // 542
// C_PUSH constant:23
@23 // 543
D=A // 544
@0 // 545
A=M // 546
M=D // 547
@0 // 548
M=M+1 // 549
// C_PUSH constant:15
@15 // 550
D=A // 551
@0 // 552
A=M // 553
M=D // 554
@0 // 555
M=M+1 // 556
// call Class2.set 2
@return.3 // 557
D=A // 558
@0 // 559
A=M // 560
M=D // 561
@0 // 562
M=M+1 // 563
@LCL // 564
D=M // 565
@0 // 566
A=M // 567
M=D // 568
@0 // 569
M=M+1 // 570
@ARG // 571
D=M // 572
@0 // 573
A=M // 574
M=D // 575
@0 // 576
M=M+1 // 577
@THIS // 578
D=M // 579
@0 // 580
A=M // 581
M=D // 582
@0 // 583
M=M+1 // 584
@THAT // 585
D=M // 586
@0 // 587
A=M // 588
M=D // 589
@0 // 590
M=M+1 // 591
@0 // 592
D=M // 593
@2 // 594
D=D-A // 595
@5 // 596
D=D-A // 597
@ARG // 598
M=D // 599
@0 // 600
D=M // 601
@LCL // 602
M=D // 603
@Class2.set // 604
0;JMP // 605
(return.3)
// C_POP temp:0
@R5 // 606
D=A // 607
@0 // 608
D=D+A // 609
@R13 // 610
M=D // 611
@0 // 612
M=M-1 // 613
A=M // 614
D=M // 615
@R13 // 616
A=M // 617
M=D // 618
// call Class1.get 0
@return.4 // 619
D=A // 620
@0 // 621
A=M // 622
M=D // 623
@0 // 624
M=M+1 // 625
@LCL // 626
D=M // 627
@0 // 628
A=M // 629
M=D // 630
@0 // 631
M=M+1 // 632
@ARG // 633
D=M // 634
@0 // 635
A=M // 636
M=D // 637
@0 // 638
M=M+1 // 639
@THIS // 640
D=M // 641
@0 // 642
A=M // 643
M=D // 644
@0 // 645
M=M+1 // 646
@THAT // 647
D=M // 648
@0 // 649
A=M // 650
M=D // 651
@0 // 652
M=M+1 // 653
@0 // 654
D=M // 655
@0 // 656
D=D-A // 657
@5 // 658
D=D-A // 659
@ARG // 660
M=D // 661
@0 // 662
D=M // 663
@LCL // 664
M=D // 665
@Class1.get // 666
0;JMP // 667
(return.4)
// call Class2.get 0
@return.5 // 668
D=A // 669
@0 // 670
A=M // 671
M=D // 672
@0 // 673
M=M+1 // 674
@LCL // 675
D=M // 676
@0 // 677
A=M // 678
M=D // 679
@0 // 680
M=M+1 // 681
@ARG // 682
D=M // 683
@0 // 684
A=M // 685
M=D // 686
@0 // 687
M=M+1 // 688
@THIS // 689
D=M // 690
@0 // 691
A=M // 692
M=D // 693
@0 // 694
M=M+1 // 695
@THAT // 696
D=M // 697
@0 // 698
A=M // 699
M=D // 700
@0 // 701
M=M+1 // 702
@0 // 703
D=M // 704
@0 // 705
D=D-A // 706
@5 // 707
D=D-A // 708
@ARG // 709
M=D // 710
@0 // 711
D=M // 712
@LCL // 713
M=D // 714
@Class2.get // 715
0;JMP // 716
(return.5)
// label WHILE
(Sys.init$WHILE)
// goto WHILE
@Sys.init$WHILE // 717
0;JMP // 718