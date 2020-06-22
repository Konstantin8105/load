package snow

// enum SNiP2_01_07_snow
// {
//     SNiP2_01_07_snow_I,
//     SNiP2_01_07_snow_II,
//     SNiP2_01_07_snow_III,
//     SNiP2_01_07_snow_IV,
//     SNiP2_01_07_snow_V,
//     SNiP2_01_07_snow_VI,
//     SNiP2_01_07_snow_VII,
//     SNiP2_01_07_snow_VIII
// };
//
// void Printf(SNiP2_01_07_snow s)
// {
//     switch(s)
//     {
//         case SNiP2_01_07_snow_I   : printf("Snow region: I\n"   );   break;
//         case SNiP2_01_07_snow_II  : printf("Snow region: II\n"  );   break;
//         case SNiP2_01_07_snow_III : printf("Snow region: III\n" );   break;
//         case SNiP2_01_07_snow_IV  : printf("Snow region: IV\n"  );   break;
//         case SNiP2_01_07_snow_V   : printf("Snow region: V\n"   );   break;
//         case SNiP2_01_07_snow_VI  : printf("Snow region: VI\n"  );   break;
//         case SNiP2_01_07_snow_VII : printf("Snow region: VII\n" );   break;
//         case SNiP2_01_07_snow_VIII: printf("Snow region: VIII\n");   break;
//         default:
//             printf("Snow region: WARNING\n");
//             FATAL();
//     }
// }
//
// void SNiP2_01_07_table4(SNiP2_01_07_snow region, double &Sg, bool OUT = true)
// {
//     switch(region)
//     {
//         case (SNiP2_01_07_snow_I   ): Sg =  80 ;   break;
//         case (SNiP2_01_07_snow_II  ): Sg = 120 ;   break;
//         case (SNiP2_01_07_snow_III ): Sg = 180 ;   break;
//         case (SNiP2_01_07_snow_IV  ): Sg = 240 ;   break;
//         case (SNiP2_01_07_snow_V   ): Sg = 320 ;   break;
//         case (SNiP2_01_07_snow_VI  ): Sg = 400 ;   break;
//         case (SNiP2_01_07_snow_VII ): Sg = 480 ;   break;
//         case (SNiP2_01_07_snow_VIII): Sg = 560 ;   break;
//         default:
//             printf("Snow region: WARNING\n");
//             FATAL();
//     }
//     if(OUT)
//     {
//         printf("Sg = %.0f kg/m2 by table 4 SNiP 2.01.07\n",Sg);
//     }
// }
