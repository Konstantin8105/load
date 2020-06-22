package seismic

// enum SNiP_II_7_Table3_K1
// {
//     SNiP_II_7_Table3_K1_row1,
//     SNiP_II_7_Table3_K1_row2,
//     SNiP_II_7_Table3_K1_row3
// };
//
// enum SNiP_II_7_Table6_K_Ksi
// {
//     SNiP_II_7_Table6_K_Ksi_row1,
//     SNiP_II_7_Table6_K_Ksi_row2,
//     SNiP_II_7_Table6_K_Ksi_row3
// };
//
// enum SNiP_II_7_region
// {
//     SNiP_II_7_region_6,
//     SNiP_II_7_region_7,
//     SNiP_II_7_region_8,
//     SNiP_II_7_region_9
// };
//
// enum SNiP_II_7_ground
// {
//     SNiP_II_7_ground_I,
//     SNiP_II_7_ground_II,
//     SNiP_II_7_ground_III
// };
//
// void SNiP_II_7_Table3(SNiP_II_7_Table3_K1 TK1, float &k1, bool OUT = true)
// {
//     switch(TK1)
//     {
//         case(SNiP_II_7_Table3_K1_row1): k1 = 1.00; break;
//         case(SNiP_II_7_Table3_K1_row2): k1 = 0.25; break;
//         case(SNiP_II_7_Table3_K1_row3): k1 = 0.12; WARNING(); break;
//         default:
//             print_name("Error in SNiP_II_7_Table3");
//             FATAL();
//     }
//     if(OUT)
//     {
//         printf("k1 = %.2f by Table 3 SNiP II-7\n",k1);
//     }
// }
//
// void SNiP_II_7_Formula_1(SNiP_II_7_Table3_K1 TK1, float Soi, float &Ski, bool OUT = true)
// {
//     float k1;
//     SNiP_II_7_Table3(TK1, k1, OUT);
//     Ski = Soi*k1;
//     if(OUT)
//     {
//         printf("Ski = Soi*K1 = %.f*%.2f = %.f by formula 1 SNiP II-7\n",Soi,k1,Ski);
//     }
// }
//
// void SNiP_II_7_Table6(SNiP_II_7_Table6_K_Ksi TKsi, float &ksi, bool OUT = true)
// {
//     switch(TKsi)
//     {
//         case(SNiP_II_7_Table6_K_Ksi_row1): ksi = 1.5; break;
//         case(SNiP_II_7_Table6_K_Ksi_row2): ksi = 1.3; break;
//         case(SNiP_II_7_Table6_K_Ksi_row3): ksi = 1.0; break;
//         default:
//             print_name("Error in SNiP_II_7_Table6");
//             FATAL();
//     }
//     if(OUT)
//     {
//         printf("ksi = %.1f by Table 6 SNiP II-7\n",ksi);
//     }
// }
//
// void SNiP_II_7_FactorA(SNiP_II_7_region Seismic_region, float &A, bool OUT = true)
// {
//     switch(Seismic_region)
//     {
//         case(SNiP_II_7_region_7): A = 0.1; break;
//         case(SNiP_II_7_region_8): A = 0.2; break;
//         case(SNiP_II_7_region_9): A = 0.4; break;
//         default:
//             print_name("Error in SNiP_II_7_FactorA");
//             FATAL();
//     }
//     if(OUT)
//     {
//         printf("A = %.1f by SNiP II-7\n",A);
//     }
// }
//
// void SNiP_II_7_Formula_2(float Qk,
//                          SNiP_II_7_region Seismic_region,
//                          float betta,
//                          SNiP_II_7_Table6_K_Ksi TKsi,
//                          float nu,
//                          float &Soi, bool OUT = true)
// {
//     if(nu > 1.1 || nu < -1.1)
//     {
//         printf("Error in SNiP_II_7_Formula_2");
//         FATAL();
//     }
//     float ksi;
//     SNiP_II_7_Table6(TKsi, ksi, OUT);
//     float A;
//     SNiP_II_7_FactorA(Seismic_region, A, OUT);
//     Soi = Qk * A * betta * ksi * nu;
//     if(OUT)
//     {
//         printf("Soi = Qk*A*betta*Ksi*nu = %.1f*%.1f*%.1f*%.1f*%.3f = %.1f by formula 2 SNiP II-7\n",
//                Qk,A,betta,ksi,nu,Soi);
//     }
// }
//
// void SNiP_II_7_Punkt_2_6(float Period, SNiP_II_7_ground ground,float &betta, bool OUT = true)
// {
//     if(ground == SNiP_II_7_ground_I || ground == SNiP_II_7_ground_II)
//     {
//                               if(Period <= 0.1) betta = 1.+15.*Period;
//         else if(0.1 <= Period && Period <  0.4) betta = 2.5;
//         else if(0.4 <= Period) betta = 2.5*sqrt(0.4/Period);
//     }
//     else
//     {
//                               if(Period <= 0.1) betta = 1.+15.*Period;
//         else if(0.1 <= Period && Period <  0.8) betta = 2.5;
//         else if(0.8 <= Period) betta = 2.5*sqrt(0.8/Period);
//     }
//     betta = max (betta, 0.8);
//     if(OUT)
//     {
//         printf("betta = %.2f for Period = %.2fsec by Formula 3 SNiP II-7-81\n",betta,Period);
//     }
// }
//
// void SNiP_II_7_Formula_8(Array<double> N, double &Np, bool OUT = true)
// {
//     Np = 0;
//     for(unsigned i=0;i<N.GetSize();i++)
//         Np += pow(N[i],2.0);
//     Np = sqrt(Np);
//     if(OUT)
//     {
//         printf("Np = %f by Formula 8 SNiP II-7-81\n",Np);
//     }
// }
//
// void SNiP_II_7_Formula_9actual(Array <double> FREQUENCY, Array<double> N, double &Np, bool OUT = true)
// {
//     bool DEBUG = false;//true;//
//     if(FREQUENCY.GetSize() != N.GetSize())
//     {
//         print_name("Error in SNiP_II_7_Formula_9actual");
//         FATAL();
//     }
//     Np = 0;
//     for(unsigned i=0;i<N.GetSize();i++)
//         Np += pow(N[i],2.0);
//     for(unsigned i=1;i<N.GetSize();i++)
//     {
//         float ro = 0;
//         if(FREQUENCY[i]/FREQUENCY[i-1] >= 0.9)
//             ro = 2;
//         Np += ro*fabs(N[i]*N[i-1]);
//     }
//     Np = sqrt(Np);
//     if(OUT)
//     {
//         if(DEBUG) Printf(N);
//         printf("Np = %f by Formula 8 SNiP II-7-81\n",Np);
//     }
// }
