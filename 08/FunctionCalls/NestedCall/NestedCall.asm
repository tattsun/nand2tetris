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
// function Sys.init 0
(Sys.init)
// C_PUSH constant:4000
@4000 // 53
D=A // 54
@0 // 55
A=M // 56
M=D // 57
@0 // 58
M=M+1 // 59
// C_POP pointer:0
@THIS // 60
D=A // 61
@0 // 62
D=D+A // 63
@R13 // 64
M=D // 65
@0 // 66
M=M-1 // 67
A=M // 68
D=M // 69
@R13 // 70
A=M // 71
M=D // 72
// C_PUSH constant:5000
@5000 // 73
D=A // 74
@0 // 75
A=M // 76
M=D // 77
@0 // 78
M=M+1 // 79
// C_POP pointer:1
@THAT // 80
D=A // 81
@0 // 82
D=D+A // 83
@R13 // 84
M=D // 85
@0 // 86
M=M-1 // 87
A=M // 88
D=M // 89
@R13 // 90
A=M // 91
M=D // 92
// call Sys.main 0
@return.2 // 93
D=A // 94
@0 // 95
A=M // 96
M=D // 97
@0 // 98
M=M+1 // 99
@LCL // 100
D=M // 101
@0 // 102
A=M // 103
M=D // 104
@0 // 105
M=M+1 // 106
@ARG // 107
D=M // 108
@0 // 109
A=M // 110
M=D // 111
@0 // 112
M=M+1 // 113
@THIS // 114
D=M // 115
@0 // 116
A=M // 117
M=D // 118
@0 // 119
M=M+1 // 120
@THAT // 121
D=M // 122
@0 // 123
A=M // 124
M=D // 125
@0 // 126
M=M+1 // 127
@0 // 128
D=M // 129
@0 // 130
D=D-A // 131
@5 // 132
D=D-A // 133
@ARG // 134
M=D // 135
@0 // 136
D=M // 137
@LCL // 138
M=D // 139
@Sys.main // 140
0;JMP // 141
(return.2)
// C_POP temp:1
@R5 // 142
D=A // 143
@1 // 144
D=D+A // 145
@R13 // 146
M=D // 147
@0 // 148
M=M-1 // 149
A=M // 150
D=M // 151
@R13 // 152
A=M // 153
M=D // 154
// label LOOP
(Sys.init$LOOP)
// goto LOOP
@Sys.init$LOOP // 155
0;JMP // 156
// function Sys.main 5
(Sys.main)
@0 // 157
D=A // 158
@0 // 159
A=M // 160
M=D // 161
@0 // 162
M=M+1 // 163
@0 // 164
D=A // 165
@0 // 166
A=M // 167
M=D // 168
@0 // 169
M=M+1 // 170
@0 // 171
D=A // 172
@0 // 173
A=M // 174
M=D // 175
@0 // 176
M=M+1 // 177
@0 // 178
D=A // 179
@0 // 180
A=M // 181
M=D // 182
@0 // 183
M=M+1 // 184
@0 // 185
D=A // 186
@0 // 187
A=M // 188
M=D // 189
@0 // 190
M=M+1 // 191
// C_PUSH constant:4001
@4001 // 192
D=A // 193
@0 // 194
A=M // 195
M=D // 196
@0 // 197
M=M+1 // 198
// C_POP pointer:0
@THIS // 199
D=A // 200
@0 // 201
D=D+A // 202
@R13 // 203
M=D // 204
@0 // 205
M=M-1 // 206
A=M // 207
D=M // 208
@R13 // 209
A=M // 210
M=D // 211
// C_PUSH constant:5001
@5001 // 212
D=A // 213
@0 // 214
A=M // 215
M=D // 216
@0 // 217
M=M+1 // 218
// C_POP pointer:1
@THAT // 219
D=A // 220
@0 // 221
D=D+A // 222
@R13 // 223
M=D // 224
@0 // 225
M=M-1 // 226
A=M // 227
D=M // 228
@R13 // 229
A=M // 230
M=D // 231
// C_PUSH constant:200
@200 // 232
D=A // 233
@0 // 234
A=M // 235
M=D // 236
@0 // 237
M=M+1 // 238
// C_POP local:1
@LCL // 239
D=M // 240
@1 // 241
D=D+A // 242
@R13 // 243
M=D // 244
@0 // 245
M=M-1 // 246
A=M // 247
D=M // 248
@R13 // 249
A=M // 250
M=D // 251
// C_PUSH constant:40
@40 // 252
D=A // 253
@0 // 254
A=M // 255
M=D // 256
@0 // 257
M=M+1 // 258
// C_POP local:2
@LCL // 259
D=M // 260
@2 // 261
D=D+A // 262
@R13 // 263
M=D // 264
@0 // 265
M=M-1 // 266
A=M // 267
D=M // 268
@R13 // 269
A=M // 270
M=D // 271
// C_PUSH constant:6
@6 // 272
D=A // 273
@0 // 274
A=M // 275
M=D // 276
@0 // 277
M=M+1 // 278
// C_POP local:3
@LCL // 279
D=M // 280
@3 // 281
D=D+A // 282
@R13 // 283
M=D // 284
@0 // 285
M=M-1 // 286
A=M // 287
D=M // 288
@R13 // 289
A=M // 290
M=D // 291
// C_PUSH constant:123
@123 // 292
D=A // 293
@0 // 294
A=M // 295
M=D // 296
@0 // 297
M=M+1 // 298
// call Sys.add12 1
@return.3 // 299
D=A // 300
@0 // 301
A=M // 302
M=D // 303
@0 // 304
M=M+1 // 305
@LCL // 306
D=M // 307
@0 // 308
A=M // 309
M=D // 310
@0 // 311
M=M+1 // 312
@ARG // 313
D=M // 314
@0 // 315
A=M // 316
M=D // 317
@0 // 318
M=M+1 // 319
@THIS // 320
D=M // 321
@0 // 322
A=M // 323
M=D // 324
@0 // 325
M=M+1 // 326
@THAT // 327
D=M // 328
@0 // 329
A=M // 330
M=D // 331
@0 // 332
M=M+1 // 333
@0 // 334
D=M // 335
@1 // 336
D=D-A // 337
@5 // 338
D=D-A // 339
@ARG // 340
M=D // 341
@0 // 342
D=M // 343
@LCL // 344
M=D // 345
@Sys.add12 // 346
0;JMP // 347
(return.3)
// C_POP temp:0
@R5 // 348
D=A // 349
@0 // 350
D=D+A // 351
@R13 // 352
M=D // 353
@0 // 354
M=M-1 // 355
A=M // 356
D=M // 357
@R13 // 358
A=M // 359
M=D // 360
// C_PUSH local:0
@LCL // 361
D=M // 362
@0 // 363
D=D+A // 364
A=D // 365
D=M // 366
@SP // 367
A=M // 368
M=D // 369
@0 // 370
M=M+1 // 371
// C_PUSH local:1
@LCL // 372
D=M // 373
@1 // 374
D=D+A // 375
A=D // 376
D=M // 377
@SP // 378
A=M // 379
M=D // 380
@0 // 381
M=M+1 // 382
// C_PUSH local:2
@LCL // 383
D=M // 384
@2 // 385
D=D+A // 386
A=D // 387
D=M // 388
@SP // 389
A=M // 390
M=D // 391
@0 // 392
M=M+1 // 393
// C_PUSH local:3
@LCL // 394
D=M // 395
@3 // 396
D=D+A // 397
A=D // 398
D=M // 399
@SP // 400
A=M // 401
M=D // 402
@0 // 403
M=M+1 // 404
// C_PUSH local:4
@LCL // 405
D=M // 406
@4 // 407
D=D+A // 408
A=D // 409
D=M // 410
@SP // 411
A=M // 412
M=D // 413
@0 // 414
M=M+1 // 415
// add
@0 // 416
M=M-1 // 417
A=M // 418
D=M // 419
@0 // 420
M=M-1 // 421
A=M // 422
M=D+M // 423
@0 // 424
M=M+1 // 425
// add
@0 // 426
M=M-1 // 427
A=M // 428
D=M // 429
@0 // 430
M=M-1 // 431
A=M // 432
M=D+M // 433
@0 // 434
M=M+1 // 435
// add
@0 // 436
M=M-1 // 437
A=M // 438
D=M // 439
@0 // 440
M=M-1 // 441
A=M // 442
M=D+M // 443
@0 // 444
M=M+1 // 445
// add
@0 // 446
M=M-1 // 447
A=M // 448
D=M // 449
@0 // 450
M=M-1 // 451
A=M // 452
M=D+M // 453
@0 // 454
M=M+1 // 455
// return
@LCL // 456
D=M // 457
@R13 // 458
M=D // 459
@R13 // 460
D=M // 461
@5 // 462
D=D-A // 463
A=D // 464
D=M // 465
@R14 // 466
M=D // 467
@0 // 468
M=M-1 // 469
A=M // 470
D=M // 471
@ARG // 472
A=M // 473
M=D // 474
@ARG // 475
D=M // 476
@1 // 477
D=D+A // 478
@SP // 479
M=D // 480
@R13 // 481
D=M // 482
@1 // 483
D=D-A // 484
A=D // 485
D=M // 486
@THAT // 487
M=D // 488
@R13 // 489
D=M // 490
@2 // 491
D=D-A // 492
A=D // 493
D=M // 494
@THIS // 495
M=D // 496
@R13 // 497
D=M // 498
@3 // 499
D=D-A // 500
A=D // 501
D=M // 502
@ARG // 503
M=D // 504
@R13 // 505
D=M // 506
@4 // 507
D=D-A // 508
A=D // 509
D=M // 510
@LCL // 511
M=D // 512
@R14 // 513
A=M // 514
0;JMP // 515
// function Sys.add12 0
(Sys.add12)
// C_PUSH constant:4002
@4002 // 516
D=A // 517
@0 // 518
A=M // 519
M=D // 520
@0 // 521
M=M+1 // 522
// C_POP pointer:0
@THIS // 523
D=A // 524
@0 // 525
D=D+A // 526
@R13 // 527
M=D // 528
@0 // 529
M=M-1 // 530
A=M // 531
D=M // 532
@R13 // 533
A=M // 534
M=D // 535
// C_PUSH constant:5002
@5002 // 536
D=A // 537
@0 // 538
A=M // 539
M=D // 540
@0 // 541
M=M+1 // 542
// C_POP pointer:1
@THAT // 543
D=A // 544
@0 // 545
D=D+A // 546
@R13 // 547
M=D // 548
@0 // 549
M=M-1 // 550
A=M // 551
D=M // 552
@R13 // 553
A=M // 554
M=D // 555
// C_PUSH argument:0
@ARG // 556
D=M // 557
@0 // 558
D=D+A // 559
A=D // 560
D=M // 561
@SP // 562
A=M // 563
M=D // 564
@0 // 565
M=M+1 // 566
// C_PUSH constant:12
@12 // 567
D=A // 568
@0 // 569
A=M // 570
M=D // 571
@0 // 572
M=M+1 // 573
// add
@0 // 574
M=M-1 // 575
A=M // 576
D=M // 577
@0 // 578
M=M-1 // 579
A=M // 580
M=D+M // 581
@0 // 582
M=M+1 // 583
// return
@LCL // 584
D=M // 585
@R13 // 586
M=D // 587
@R13 // 588
D=M // 589
@5 // 590
D=D-A // 591
A=D // 592
D=M // 593
@R14 // 594
M=D // 595
@0 // 596
M=M-1 // 597
A=M // 598
D=M // 599
@ARG // 600
A=M // 601
M=D // 602
@ARG // 603
D=M // 604
@1 // 605
D=D+A // 606
@SP // 607
M=D // 608
@R13 // 609
D=M // 610
@1 // 611
D=D-A // 612
A=D // 613
D=M // 614
@THAT // 615
M=D // 616
@R13 // 617
D=M // 618
@2 // 619
D=D-A // 620
A=D // 621
D=M // 622
@THIS // 623
M=D // 624
@R13 // 625
D=M // 626
@3 // 627
D=D-A // 628
A=D // 629
D=M // 630
@ARG // 631
M=D // 632
@R13 // 633
D=M // 634
@4 // 635
D=D-A // 636
A=D // 637
D=M // 638
@LCL // 639
M=D // 640
@R14 // 641
A=M // 642
0;JMP // 643
