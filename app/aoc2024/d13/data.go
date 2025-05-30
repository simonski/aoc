package d13

const TEST_DATA = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

const REAL_DATA = `Button A: X+17, Y+68
Button B: X+62, Y+15
Prize: X=4507, Y=17764

Button A: X+19, Y+95
Button B: X+59, Y+23
Prize: X=6550, Y=10174

Button A: X+20, Y+58
Button B: X+60, Y+11
Prize: X=3140, Y=4691

Button A: X+92, Y+31
Button B: X+49, Y+87
Prize: X=6599, Y=2717

Button A: X+46, Y+72
Button B: X+29, Y+12
Prize: X=18754, Y=13304

Button A: X+42, Y+33
Button B: X+19, Y+81
Prize: X=3531, Y=6144

Button A: X+80, Y+13
Button B: X+20, Y+75
Prize: X=1860, Y=3818

Button A: X+79, Y+14
Button B: X+56, Y+66
Prize: X=6740, Y=3830

Button A: X+18, Y+85
Button B: X+70, Y+36
Prize: X=5358, Y=5861

Button A: X+42, Y+44
Button B: X+52, Y+14
Prize: X=6310, Y=4870

Button A: X+56, Y+90
Button B: X+88, Y+41
Prize: X=9208, Y=8572

Button A: X+95, Y+67
Button B: X+13, Y+71
Prize: X=3065, Y=6799

Button A: X+70, Y+22
Button B: X+13, Y+57
Prize: X=17469, Y=5153

Button A: X+36, Y+65
Button B: X+79, Y+27
Prize: X=5764, Y=4394

Button A: X+20, Y+46
Button B: X+67, Y+29
Prize: X=4637, Y=10139

Button A: X+47, Y+72
Button B: X+42, Y+20
Prize: X=1495, Y=10984

Button A: X+37, Y+96
Button B: X+70, Y+53
Prize: X=3563, Y=6029

Button A: X+20, Y+59
Button B: X+62, Y+13
Prize: X=5550, Y=18133

Button A: X+16, Y+52
Button B: X+70, Y+61
Prize: X=2400, Y=5136

Button A: X+70, Y+32
Button B: X+22, Y+53
Prize: X=320, Y=18668

Button A: X+47, Y+78
Button B: X+78, Y+17
Prize: X=6317, Y=6323

Button A: X+48, Y+33
Button B: X+20, Y+89
Prize: X=4296, Y=7017

Button A: X+17, Y+53
Button B: X+43, Y+14
Prize: X=19634, Y=5590

Button A: X+41, Y+95
Button B: X+78, Y+30
Prize: X=6816, Y=6900

Button A: X+14, Y+75
Button B: X+65, Y+12
Prize: X=10103, Y=12239

Button A: X+84, Y+69
Button B: X+23, Y+95
Prize: X=6178, Y=11620

Button A: X+57, Y+88
Button B: X+70, Y+33
Prize: X=4926, Y=6479

Button A: X+84, Y+34
Button B: X+29, Y+78
Prize: X=6138, Y=2882

Button A: X+32, Y+60
Button B: X+62, Y+29
Prize: X=13264, Y=10848

Button A: X+81, Y+69
Button B: X+13, Y+61
Prize: X=7893, Y=11217

Button A: X+98, Y+12
Button B: X+85, Y+85
Prize: X=8352, Y=3708

Button A: X+22, Y+91
Button B: X+73, Y+39
Prize: X=1457, Y=4186

Button A: X+38, Y+56
Button B: X+97, Y+16
Prize: X=9339, Y=3480

Button A: X+46, Y+13
Button B: X+36, Y+71
Prize: X=14918, Y=17793

Button A: X+27, Y+74
Button B: X+64, Y+13
Prize: X=11637, Y=3584

Button A: X+18, Y+58
Button B: X+35, Y+16
Prize: X=13544, Y=2948

Button A: X+32, Y+99
Button B: X+99, Y+16
Prize: X=7159, Y=4441

Button A: X+86, Y+49
Button B: X+12, Y+49
Prize: X=13188, Y=6269

Button A: X+19, Y+72
Button B: X+37, Y+27
Prize: X=1001, Y=3114

Button A: X+80, Y+39
Button B: X+36, Y+55
Prize: X=3848, Y=4797

Button A: X+32, Y+77
Button B: X+82, Y+21
Prize: X=4120, Y=4977

Button A: X+41, Y+64
Button B: X+41, Y+16
Prize: X=1829, Y=5360

Button A: X+67, Y+45
Button B: X+13, Y+29
Prize: X=19806, Y=5416

Button A: X+84, Y+13
Button B: X+37, Y+55
Prize: X=2852, Y=3792

Button A: X+70, Y+37
Button B: X+21, Y+52
Prize: X=8089, Y=10953

Button A: X+16, Y+57
Button B: X+81, Y+38
Prize: X=18875, Y=17714

Button A: X+91, Y+87
Button B: X+24, Y+97
Prize: X=3019, Y=5182

Button A: X+48, Y+60
Button B: X+75, Y+17
Prize: X=3267, Y=937

Button A: X+78, Y+26
Button B: X+54, Y+95
Prize: X=10638, Y=9706

Button A: X+21, Y+90
Button B: X+65, Y+18
Prize: X=5264, Y=4320

Button A: X+14, Y+58
Button B: X+41, Y+33
Prize: X=931, Y=983

Button A: X+22, Y+45
Button B: X+55, Y+23
Prize: X=18987, Y=18336

Button A: X+24, Y+61
Button B: X+75, Y+29
Prize: X=5127, Y=5758

Button A: X+49, Y+84
Button B: X+38, Y+11
Prize: X=3463, Y=3446

Button A: X+32, Y+75
Button B: X+54, Y+11
Prize: X=10780, Y=933

Button A: X+18, Y+84
Button B: X+75, Y+31
Prize: X=2577, Y=2137

Button A: X+16, Y+53
Button B: X+66, Y+24
Prize: X=12502, Y=7650

Button A: X+46, Y+75
Button B: X+46, Y+20
Prize: X=7598, Y=5780

Button A: X+97, Y+57
Button B: X+40, Y+75
Prize: X=6896, Y=6936

Button A: X+87, Y+49
Button B: X+17, Y+61
Prize: X=2245, Y=5687

Button A: X+89, Y+43
Button B: X+38, Y+88
Prize: X=8665, Y=7181

Button A: X+39, Y+94
Button B: X+34, Y+26
Prize: X=2237, Y=5112

Button A: X+57, Y+12
Button B: X+40, Y+71
Prize: X=5110, Y=6395

Button A: X+92, Y+13
Button B: X+29, Y+88
Prize: X=6278, Y=4411

Button A: X+94, Y+81
Button B: X+13, Y+64
Prize: X=3211, Y=5882

Button A: X+49, Y+34
Button B: X+14, Y+65
Prize: X=1939, Y=3778

Button A: X+32, Y+73
Button B: X+50, Y+12
Prize: X=12774, Y=12247

Button A: X+34, Y+97
Button B: X+78, Y+59
Prize: X=1364, Y=1602

Button A: X+13, Y+44
Button B: X+70, Y+41
Prize: X=13681, Y=7491

Button A: X+72, Y+44
Button B: X+13, Y+56
Prize: X=1078, Y=1716

Button A: X+56, Y+30
Button B: X+27, Y+58
Prize: X=3425, Y=5100

Button A: X+99, Y+18
Button B: X+40, Y+83
Prize: X=9389, Y=6175

Button A: X+44, Y+14
Button B: X+17, Y+66
Prize: X=7836, Y=16974

Button A: X+16, Y+58
Button B: X+79, Y+29
Prize: X=1201, Y=18041

Button A: X+60, Y+19
Button B: X+52, Y+73
Prize: X=5560, Y=5718

Button A: X+13, Y+51
Button B: X+69, Y+26
Prize: X=7423, Y=7244

Button A: X+40, Y+14
Button B: X+42, Y+60
Prize: X=6320, Y=5836

Button A: X+73, Y+29
Button B: X+20, Y+54
Prize: X=12274, Y=12544

Button A: X+12, Y+64
Button B: X+89, Y+95
Prize: X=8072, Y=9640

Button A: X+26, Y+49
Button B: X+70, Y+26
Prize: X=2276, Y=3442

Button A: X+68, Y+47
Button B: X+29, Y+69
Prize: X=7480, Y=8499

Button A: X+13, Y+28
Button B: X+86, Y+21
Prize: X=1826, Y=1141

Button A: X+97, Y+85
Button B: X+25, Y+88
Prize: X=5021, Y=6647

Button A: X+24, Y+50
Button B: X+85, Y+17
Prize: X=1952, Y=2786

Button A: X+21, Y+71
Button B: X+59, Y+21
Prize: X=18503, Y=12693

Button A: X+94, Y+58
Button B: X+42, Y+85
Prize: X=12332, Y=12395

Button A: X+93, Y+49
Button B: X+13, Y+40
Prize: X=9332, Y=7370

Button A: X+58, Y+44
Button B: X+32, Y+84
Prize: X=6696, Y=10216

Button A: X+18, Y+77
Button B: X+56, Y+51
Prize: X=2952, Y=5840

Button A: X+13, Y+76
Button B: X+46, Y+11
Prize: X=14558, Y=6711

Button A: X+46, Y+20
Button B: X+45, Y+70
Prize: X=6149, Y=5750

Button A: X+92, Y+39
Button B: X+54, Y+97
Prize: X=10056, Y=10488

Button A: X+25, Y+48
Button B: X+80, Y+50
Prize: X=2260, Y=1542

Button A: X+12, Y+52
Button B: X+39, Y+20
Prize: X=2225, Y=10452

Button A: X+25, Y+62
Button B: X+42, Y+15
Prize: X=2596, Y=12552

Button A: X+11, Y+39
Button B: X+67, Y+17
Prize: X=2192, Y=12338

Button A: X+90, Y+59
Button B: X+19, Y+57
Prize: X=5277, Y=7602

Button A: X+50, Y+34
Button B: X+33, Y+80
Prize: X=1245, Y=1710

Button A: X+11, Y+63
Button B: X+83, Y+12
Prize: X=4173, Y=15554

Button A: X+52, Y+97
Button B: X+91, Y+13
Prize: X=13091, Y=10469

Button A: X+23, Y+57
Button B: X+69, Y+36
Prize: X=8771, Y=6614

Button A: X+21, Y+40
Button B: X+67, Y+40
Prize: X=15512, Y=5120

Button A: X+20, Y+47
Button B: X+66, Y+22
Prize: X=9426, Y=16922

Button A: X+15, Y+37
Button B: X+49, Y+14
Prize: X=2334, Y=3915

Button A: X+25, Y+51
Button B: X+44, Y+17
Prize: X=4718, Y=4386

Button A: X+77, Y+30
Button B: X+16, Y+78
Prize: X=7755, Y=7758

Button A: X+55, Y+87
Button B: X+64, Y+15
Prize: X=10939, Y=9456

Button A: X+64, Y+17
Button B: X+24, Y+78
Prize: X=312, Y=3577

Button A: X+15, Y+68
Button B: X+73, Y+12
Prize: X=708, Y=6800

Button A: X+71, Y+66
Button B: X+12, Y+42
Prize: X=7143, Y=8028

Button A: X+70, Y+24
Button B: X+22, Y+66
Prize: X=1930, Y=10706

Button A: X+52, Y+28
Button B: X+29, Y+47
Prize: X=19597, Y=3007

Button A: X+11, Y+24
Button B: X+51, Y+24
Prize: X=2230, Y=1200

Button A: X+23, Y+80
Button B: X+55, Y+15
Prize: X=2251, Y=1130

Button A: X+20, Y+42
Button B: X+75, Y+48
Prize: X=15760, Y=14762

Button A: X+63, Y+12
Button B: X+32, Y+77
Prize: X=1622, Y=1319

Button A: X+11, Y+20
Button B: X+54, Y+22
Prize: X=10462, Y=1050

Button A: X+24, Y+72
Button B: X+53, Y+12
Prize: X=5360, Y=10304

Button A: X+99, Y+52
Button B: X+12, Y+52
Prize: X=8937, Y=5928

Button A: X+62, Y+25
Button B: X+57, Y+95
Prize: X=10249, Y=10110

Button A: X+40, Y+99
Button B: X+91, Y+13
Prize: X=3050, Y=1182

Button A: X+26, Y+60
Button B: X+71, Y+38
Prize: X=9324, Y=9448

Button A: X+60, Y+88
Button B: X+73, Y+32
Prize: X=8559, Y=7824

Button A: X+59, Y+21
Button B: X+59, Y+75
Prize: X=4602, Y=2070

Button A: X+14, Y+48
Button B: X+47, Y+19
Prize: X=4601, Y=3977

Button A: X+64, Y+22
Button B: X+29, Y+65
Prize: X=12970, Y=19186

Button A: X+28, Y+19
Button B: X+18, Y+43
Prize: X=2002, Y=2005

Button A: X+24, Y+17
Button B: X+12, Y+30
Prize: X=14612, Y=12593

Button A: X+34, Y+54
Button B: X+88, Y+12
Prize: X=9576, Y=2688

Button A: X+22, Y+35
Button B: X+51, Y+17
Prize: X=6367, Y=4806

Button A: X+28, Y+53
Button B: X+91, Y+44
Prize: X=5194, Y=3932

Button A: X+31, Y+11
Button B: X+30, Y+69
Prize: X=14791, Y=1482

Button A: X+46, Y+94
Button B: X+26, Y+17
Prize: X=3028, Y=3514

Button A: X+79, Y+29
Button B: X+13, Y+52
Prize: X=16594, Y=6031

Button A: X+76, Y+36
Button B: X+32, Y+80
Prize: X=5908, Y=3836

Button A: X+20, Y+41
Button B: X+59, Y+28
Prize: X=4243, Y=10632

Button A: X+72, Y+32
Button B: X+22, Y+58
Prize: X=15972, Y=11444

Button A: X+53, Y+28
Button B: X+28, Y+58
Prize: X=3946, Y=10646

Button A: X+19, Y+69
Button B: X+61, Y+22
Prize: X=9244, Y=17018

Button A: X+20, Y+71
Button B: X+64, Y+18
Prize: X=17356, Y=9567

Button A: X+28, Y+89
Button B: X+69, Y+50
Prize: X=5623, Y=5174

Button A: X+62, Y+17
Button B: X+11, Y+73
Prize: X=1851, Y=10184

Button A: X+60, Y+22
Button B: X+18, Y+38
Prize: X=3728, Y=6252

Button A: X+58, Y+32
Button B: X+28, Y+94
Prize: X=5286, Y=6844

Button A: X+53, Y+80
Button B: X+58, Y+14
Prize: X=8236, Y=5886

Button A: X+15, Y+49
Button B: X+46, Y+18
Prize: X=11211, Y=18229

Button A: X+21, Y+96
Button B: X+27, Y+29
Prize: X=2448, Y=7697

Button A: X+42, Y+85
Button B: X+93, Y+37
Prize: X=5364, Y=5412

Button A: X+56, Y+32
Button B: X+16, Y+31
Prize: X=3368, Y=3389

Button A: X+14, Y+66
Button B: X+81, Y+66
Prize: X=7042, Y=11088

Button A: X+12, Y+54
Button B: X+76, Y+37
Prize: X=5280, Y=15420

Button A: X+37, Y+24
Button B: X+12, Y+44
Prize: X=4537, Y=5804

Button A: X+55, Y+14
Button B: X+24, Y+85
Prize: X=4477, Y=3743

Button A: X+70, Y+92
Button B: X+73, Y+17
Prize: X=5044, Y=2840

Button A: X+13, Y+51
Button B: X+40, Y+38
Prize: X=2505, Y=5665

Button A: X+63, Y+24
Button B: X+22, Y+55
Prize: X=6942, Y=6141

Button A: X+30, Y+75
Button B: X+62, Y+22
Prize: X=14496, Y=6211

Button A: X+87, Y+43
Button B: X+30, Y+75
Prize: X=3153, Y=6312

Button A: X+77, Y+37
Button B: X+14, Y+45
Prize: X=11197, Y=3283

Button A: X+48, Y+25
Button B: X+49, Y+96
Prize: X=1254, Y=1076

Button A: X+14, Y+77
Button B: X+53, Y+55
Prize: X=1972, Y=5170

Button A: X+32, Y+19
Button B: X+18, Y+39
Prize: X=3720, Y=3681

Button A: X+36, Y+13
Button B: X+39, Y+75
Prize: X=5181, Y=6196

Button A: X+34, Y+23
Button B: X+11, Y+35
Prize: X=14981, Y=13421

Button A: X+83, Y+28
Button B: X+11, Y+47
Prize: X=5038, Y=15145

Button A: X+44, Y+98
Button B: X+78, Y+26
Prize: X=7832, Y=4444

Button A: X+18, Y+11
Button B: X+19, Y+44
Prize: X=15969, Y=15082

Button A: X+14, Y+81
Button B: X+39, Y+34
Prize: X=1722, Y=1914

Button A: X+29, Y+94
Button B: X+75, Y+55
Prize: X=9696, Y=14311

Button A: X+41, Y+79
Button B: X+58, Y+19
Prize: X=4374, Y=5367

Button A: X+47, Y+14
Button B: X+35, Y+66
Prize: X=17301, Y=1534

Button A: X+68, Y+29
Button B: X+21, Y+98
Prize: X=6996, Y=5833

Button A: X+16, Y+61
Button B: X+72, Y+15
Prize: X=10896, Y=9927

Button A: X+69, Y+16
Button B: X+17, Y+99
Prize: X=4411, Y=2924

Button A: X+16, Y+73
Button B: X+43, Y+42
Prize: X=2463, Y=4299

Button A: X+54, Y+16
Button B: X+21, Y+43
Prize: X=4187, Y=10273

Button A: X+49, Y+13
Button B: X+49, Y+83
Prize: X=6076, Y=4902

Button A: X+40, Y+19
Button B: X+17, Y+31
Prize: X=18121, Y=6893

Button A: X+22, Y+89
Button B: X+62, Y+33
Prize: X=3470, Y=2929

Button A: X+24, Y+51
Button B: X+51, Y+11
Prize: X=13013, Y=7760

Button A: X+40, Y+72
Button B: X+94, Y+29
Prize: X=10730, Y=8799

Button A: X+85, Y+18
Button B: X+64, Y+86
Prize: X=3858, Y=4222

Button A: X+99, Y+11
Button B: X+18, Y+71
Prize: X=1530, Y=722

Button A: X+69, Y+12
Button B: X+93, Y+98
Prize: X=11136, Y=9792

Button A: X+95, Y+75
Button B: X+21, Y+96
Prize: X=917, Y=1677

Button A: X+22, Y+57
Button B: X+46, Y+17
Prize: X=7332, Y=17866

Button A: X+55, Y+19
Button B: X+27, Y+67
Prize: X=8606, Y=17118

Button A: X+42, Y+92
Button B: X+77, Y+50
Prize: X=3416, Y=2736

Button A: X+13, Y+64
Button B: X+92, Y+86
Prize: X=5175, Y=6030

Button A: X+56, Y+18
Button B: X+20, Y+46
Prize: X=4068, Y=5212

Button A: X+78, Y+11
Button B: X+19, Y+79
Prize: X=6081, Y=10602

Button A: X+19, Y+87
Button B: X+84, Y+35
Prize: X=3342, Y=3765

Button A: X+57, Y+33
Button B: X+19, Y+43
Prize: X=8062, Y=8686

Button A: X+93, Y+60
Button B: X+14, Y+72
Prize: X=1229, Y=2556

Button A: X+55, Y+24
Button B: X+14, Y+94
Prize: X=1871, Y=6002

Button A: X+13, Y+19
Button B: X+37, Y+15
Prize: X=11560, Y=7340

Button A: X+97, Y+92
Button B: X+68, Y+14
Prize: X=12918, Y=8566

Button A: X+64, Y+16
Button B: X+19, Y+58
Prize: X=15227, Y=15866

Button A: X+12, Y+62
Button B: X+63, Y+17
Prize: X=11813, Y=12953

Button A: X+72, Y+28
Button B: X+20, Y+52
Prize: X=1284, Y=4264

Button A: X+76, Y+17
Button B: X+23, Y+38
Prize: X=8199, Y=4101

Button A: X+84, Y+38
Button B: X+17, Y+47
Prize: X=7592, Y=6422

Button A: X+76, Y+26
Button B: X+33, Y+61
Prize: X=7100, Y=3622

Button A: X+53, Y+45
Button B: X+14, Y+48
Prize: X=3495, Y=4737

Button A: X+14, Y+57
Button B: X+95, Y+13
Prize: X=2193, Y=4817

Button A: X+91, Y+24
Button B: X+39, Y+44
Prize: X=8905, Y=4776

Button A: X+36, Y+95
Button B: X+54, Y+14
Prize: X=6570, Y=8985

Button A: X+47, Y+26
Button B: X+27, Y+45
Prize: X=9856, Y=1870

Button A: X+70, Y+23
Button B: X+33, Y+80
Prize: X=6705, Y=3932

Button A: X+93, Y+22
Button B: X+72, Y+97
Prize: X=13317, Y=8748

Button A: X+68, Y+56
Button B: X+12, Y+82
Prize: X=3792, Y=5070

Button A: X+91, Y+19
Button B: X+50, Y+58
Prize: X=7149, Y=4917

Button A: X+17, Y+85
Button B: X+62, Y+56
Prize: X=6472, Y=11278

Button A: X+71, Y+18
Button B: X+16, Y+75
Prize: X=7574, Y=7170

Button A: X+13, Y+78
Button B: X+64, Y+76
Prize: X=3350, Y=6856

Button A: X+65, Y+14
Button B: X+16, Y+42
Prize: X=8942, Y=19674

Button A: X+37, Y+19
Button B: X+20, Y+33
Prize: X=15445, Y=3074

Button A: X+12, Y+40
Button B: X+76, Y+26
Prize: X=1848, Y=19728

Button A: X+86, Y+16
Button B: X+11, Y+96
Prize: X=7116, Y=9216

Button A: X+62, Y+14
Button B: X+49, Y+69
Prize: X=5138, Y=6722

Button A: X+44, Y+63
Button B: X+42, Y+17
Prize: X=10920, Y=3470

Button A: X+51, Y+26
Button B: X+35, Y+55
Prize: X=11198, Y=1063

Button A: X+85, Y+11
Button B: X+11, Y+72
Prize: X=507, Y=16765

Button A: X+50, Y+17
Button B: X+28, Y+55
Prize: X=2500, Y=4978

Button A: X+17, Y+41
Button B: X+60, Y+38
Prize: X=11731, Y=5537

Button A: X+31, Y+41
Button B: X+91, Y+36
Prize: X=9013, Y=6353

Button A: X+62, Y+21
Button B: X+15, Y+59
Prize: X=1794, Y=13288

Button A: X+41, Y+14
Button B: X+26, Y+49
Prize: X=7670, Y=3910

Button A: X+24, Y+99
Button B: X+95, Y+54
Prize: X=8639, Y=10971

Button A: X+13, Y+80
Button B: X+93, Y+52
Prize: X=4898, Y=7248

Button A: X+97, Y+73
Button B: X+31, Y+83
Prize: X=6926, Y=10702

Button A: X+21, Y+86
Button B: X+90, Y+65
Prize: X=7887, Y=8317

Button A: X+37, Y+42
Button B: X+83, Y+27
Prize: X=7887, Y=3105

Button A: X+19, Y+68
Button B: X+67, Y+25
Prize: X=3042, Y=7375

Button A: X+30, Y+51
Button B: X+80, Y+15
Prize: X=2140, Y=2670

Button A: X+87, Y+22
Button B: X+49, Y+91
Prize: X=6983, Y=3338

Button A: X+74, Y+42
Button B: X+14, Y+40
Prize: X=8560, Y=1180

Button A: X+98, Y+11
Button B: X+53, Y+87
Prize: X=4734, Y=3125

Button A: X+36, Y+60
Button B: X+94, Y+15
Prize: X=6512, Y=2070

Button A: X+36, Y+11
Button B: X+33, Y+57
Prize: X=15008, Y=16375

Button A: X+79, Y+51
Button B: X+18, Y+46
Prize: X=12284, Y=18892

Button A: X+25, Y+63
Button B: X+66, Y+18
Prize: X=12741, Y=16919

Button A: X+11, Y+54
Button B: X+76, Y+33
Prize: X=4920, Y=6683

Button A: X+23, Y+71
Button B: X+55, Y+17
Prize: X=15713, Y=18075

Button A: X+92, Y+85
Button B: X+72, Y+14
Prize: X=11088, Y=6988

Button A: X+90, Y+31
Button B: X+15, Y+41
Prize: X=8280, Y=5217

Button A: X+16, Y+52
Button B: X+72, Y+24
Prize: X=15496, Y=13792

Button A: X+17, Y+24
Button B: X+76, Y+16
Prize: X=8548, Y=3760

Button A: X+21, Y+42
Button B: X+56, Y+20
Prize: X=3955, Y=3034

Button A: X+57, Y+15
Button B: X+16, Y+20
Prize: X=3915, Y=1125

Button A: X+64, Y+16
Button B: X+25, Y+70
Prize: X=12199, Y=10666

Button A: X+13, Y+57
Button B: X+36, Y+13
Prize: X=12196, Y=1889

Button A: X+50, Y+70
Button B: X+29, Y+13
Prize: X=19005, Y=11065

Button A: X+58, Y+87
Button B: X+76, Y+34
Prize: X=6738, Y=7227

Button A: X+79, Y+62
Button B: X+24, Y+88
Prize: X=2126, Y=1876

Button A: X+29, Y+89
Button B: X+70, Y+22
Prize: X=7399, Y=4003

Button A: X+61, Y+36
Button B: X+12, Y+47
Prize: X=7624, Y=15339

Button A: X+16, Y+27
Button B: X+43, Y+14
Prize: X=15461, Y=6925

Button A: X+53, Y+14
Button B: X+17, Y+53
Prize: X=1833, Y=2085

Button A: X+12, Y+46
Button B: X+76, Y+33
Prize: X=17232, Y=7631

Button A: X+23, Y+15
Button B: X+11, Y+32
Prize: X=2079, Y=2051

Button A: X+14, Y+61
Button B: X+79, Y+31
Prize: X=19076, Y=19749

Button A: X+67, Y+97
Button B: X+95, Y+39
Prize: X=12736, Y=11344

Button A: X+14, Y+23
Button B: X+35, Y+11
Prize: X=6444, Y=17109

Button A: X+14, Y+40
Button B: X+32, Y+17
Prize: X=5020, Y=6220

Button A: X+53, Y+14
Button B: X+14, Y+42
Prize: X=17332, Y=3966

Button A: X+43, Y+16
Button B: X+12, Y+25
Prize: X=2577, Y=12758

Button A: X+26, Y+12
Button B: X+24, Y+47
Prize: X=17828, Y=3061

Button A: X+99, Y+66
Button B: X+43, Y+91
Prize: X=6135, Y=8391

Button A: X+13, Y+41
Button B: X+63, Y+12
Prize: X=6296, Y=12458

Button A: X+18, Y+42
Button B: X+88, Y+22
Prize: X=2486, Y=1034

Button A: X+35, Y+80
Button B: X+52, Y+37
Prize: X=5416, Y=8041

Button A: X+81, Y+32
Button B: X+21, Y+50
Prize: X=3750, Y=1690

Button A: X+90, Y+27
Button B: X+17, Y+82
Prize: X=3355, Y=6005

Button A: X+35, Y+14
Button B: X+35, Y+48
Prize: X=13850, Y=6312

Button A: X+94, Y+82
Button B: X+14, Y+41
Prize: X=4962, Y=6027

Button A: X+18, Y+86
Button B: X+36, Y+31
Prize: X=2826, Y=8849

Button A: X+24, Y+42
Button B: X+57, Y+26
Prize: X=4289, Y=16742

Button A: X+15, Y+59
Button B: X+76, Y+27
Prize: X=12775, Y=18008

Button A: X+64, Y+27
Button B: X+20, Y+65
Prize: X=19348, Y=11874

Button A: X+27, Y+55
Button B: X+85, Y+14
Prize: X=2653, Y=1903

Button A: X+90, Y+12
Button B: X+47, Y+90
Prize: X=681, Y=342

Button A: X+13, Y+29
Button B: X+18, Y+13
Prize: X=2527, Y=8550

Button A: X+71, Y+34
Button B: X+25, Y+59
Prize: X=10411, Y=6369

Button A: X+78, Y+26
Button B: X+61, Y+98
Prize: X=7284, Y=3360

Button A: X+49, Y+78
Button B: X+47, Y+11
Prize: X=8090, Y=8347

Button A: X+65, Y+29
Button B: X+23, Y+59
Prize: X=3886, Y=2806

Button A: X+13, Y+58
Button B: X+82, Y+33
Prize: X=17126, Y=1844

Button A: X+15, Y+46
Button B: X+52, Y+29
Prize: X=14512, Y=8557

Button A: X+51, Y+19
Button B: X+17, Y+57
Prize: X=4573, Y=2565

Button A: X+16, Y+84
Button B: X+43, Y+29
Prize: X=2326, Y=2374

Button A: X+21, Y+80
Button B: X+52, Y+12
Prize: X=2872, Y=11268

Button A: X+16, Y+56
Button B: X+42, Y+20
Prize: X=1218, Y=17444

Button A: X+43, Y+14
Button B: X+25, Y+38
Prize: X=2364, Y=11092

Button A: X+34, Y+16
Button B: X+42, Y+65
Prize: X=19300, Y=16358

Button A: X+12, Y+47
Button B: X+68, Y+29
Prize: X=17592, Y=3390

Button A: X+38, Y+11
Button B: X+55, Y+70
Prize: X=4021, Y=1867

Button A: X+33, Y+13
Button B: X+14, Y+52
Prize: X=8642, Y=16058

Button A: X+74, Y+27
Button B: X+12, Y+54
Prize: X=5318, Y=5265

Button A: X+28, Y+44
Button B: X+38, Y+11
Prize: X=11126, Y=12365

Button A: X+34, Y+56
Button B: X+44, Y+20
Prize: X=14102, Y=5556

Button A: X+75, Y+85
Button B: X+53, Y+16
Prize: X=6607, Y=5549

Button A: X+81, Y+20
Button B: X+35, Y+57
Prize: X=2908, Y=3281

Button A: X+13, Y+30
Button B: X+41, Y+19
Prize: X=10899, Y=10340

Button A: X+11, Y+25
Button B: X+64, Y+22
Prize: X=18766, Y=8966

Button A: X+57, Y+15
Button B: X+17, Y+42
Prize: X=14691, Y=8777

Button A: X+19, Y+29
Button B: X+72, Y+16
Prize: X=7014, Y=3194

Button A: X+22, Y+57
Button B: X+44, Y+21
Prize: X=15632, Y=19217

Button A: X+48, Y+94
Button B: X+86, Y+37
Prize: X=11514, Y=11115

Button A: X+11, Y+53
Button B: X+57, Y+22
Prize: X=8712, Y=5709

Button A: X+92, Y+23
Button B: X+12, Y+81
Prize: X=7868, Y=7661

Button A: X+52, Y+79
Button B: X+63, Y+22
Prize: X=5604, Y=2322

Button A: X+20, Y+51
Button B: X+95, Y+41
Prize: X=9500, Y=8125

Button A: X+24, Y+85
Button B: X+68, Y+46
Prize: X=7808, Y=9339

Button A: X+16, Y+58
Button B: X+83, Y+72
Prize: X=2007, Y=4300

Button A: X+86, Y+16
Button B: X+31, Y+75
Prize: X=7780, Y=5186

Button A: X+21, Y+59
Button B: X+89, Y+41
Prize: X=2578, Y=3062

Button A: X+62, Y+13
Button B: X+34, Y+86
Prize: X=5006, Y=6334

Button A: X+75, Y+86
Button B: X+80, Y+11
Prize: X=6665, Y=1749

Button A: X+20, Y+55
Button B: X+59, Y+47
Prize: X=2643, Y=3004`
