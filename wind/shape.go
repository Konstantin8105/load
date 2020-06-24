package wind

import "math"

// see part B.1.11
func Sphere(zone Zone, rg Region, zg, d, Δ float64) (cx, cz, ν float64, err error) {
	// TODO : add error handling
	ze := zg + d/2.0

	// коэффициент подъемной силы сферы
	if zg > d/2 {
		cz = 0.0
	} else {
		cz = 0.6
	}

	Wo := float64(rg)
	Kz, err := FactorKz(zone, ze)
	if err != nil {
		return
	}
	Re := 0.88 * d * math.Sqrt(Wo*Kz*γf) * 1e5
	cx = GraphB14(d, Δ, Re)
	ν = FactorNu(0.7*d, 0.7*d)

	if zg < d/2.0 {
		cx *= 1.6
	}
	return
}

// double SNiP2_01_07_Schema12b_Ce1(double angle, double h1, double d, )
// {
//     if(angle == 0. ) return 1.0;
//     if(angle < 0.  ) return SNiP2_01_07_Schema12b_Ce1(-angle                     ,h1,d,OUT);
//     if(angle > 360.) return SNiP2_01_07_Schema12b_Ce1(angle - int(angle/360.)*360,h1,d,OUT);
//     if(angle > 180.) return SNiP2_01_07_Schema12b_Ce1(180.-(angle-180.)          ,h1,d,OUT);
//     double Otn[30] = {
//         00.,    1.0,
//         10.,    1.0,
//         20.,    0.8,
//         30.,    0.4,
//         40.,    0.0,
//         50.,    -0.6,
//         60.,    -1.2,
//         70.,    -1.3,
//         80.,    -1.2,
//         90.,    -1.0,
//         100.,   -0.8,
//         110.,   -0.4,
//         120.,   -0.4,
//         130.,   -0.4,
//         180.1,   -0.4};
// //    if(angle < Otn[00*2+0])return Otn[00*2+1];
// //    if(angle > Otn[14*2+0])return Otn[14*2+1];
//     type_LLU i;
//     for(i=0;i<15;i++)
//     {
//         //if(angle == Otn[i*2+0]) return Otn[i*2 + 1];
//         if(angle <  Otn[i*2+0]) break;
//     }
//     double Cbetta = LinearInter(Otn[i*2+1],Otn[i*2+0],Otn[(i-1)*2+1],Otn[(i-1)*2+0],angle);
//     double k1     = SNiP2_01_07_Schema12b_K1(h1, d);
//     if(Cbetta > 0) k1 = 1;
//     double Ce1 = Cbetta * k1;
//     return Ce1;
// }
//
// double SNiP2_01_07_Schema12a_Re(double Diameter,double Wo, double Kz, )
// {
//     double Re = 0.88*Diameter*sqrt(Wo*Kz*1.4)*1e5;
//     return Re;
// }
//
// double SNiP2_01_07_Schema12b_K1(double h1, double d)
// {
//     double Otn[14]={
//         0.2,    0.8,
//         0.5,    0.9,
//         1.0,    0.95,
//         2.0,    1.0,
//         5.0,    1.1,
//         10.,    1.15,
//         25.,    1.2
//     };
//     if(h1/d < Otn[0*2+0])return Otn[0*2+1];
//     if(h1/d > Otn[6*2+0])return Otn[6*2+1];
//     type_LLU i;
//     for(i=0;i<7;i++)
//     {
//         if(h1/d == Otn[i*2+0]) return Otn[i*2 + 1];
//         if(h1/d <  Otn[i*2+0]) break;
//     }
//     double k1 = LinearInter(Otn[i*2+1],Otn[i*2+0],Otn[(i-1)*2+1],Otn[(i-1)*2+0],h1/d);
//     return k1;
// }
//
//
// double SNiP2_01_07_Cxema13_table1_K(double lambda, bool OUT = true)
// {
//     double K;
//          if(lambda > 100                ) K = 1.00;
//     else if(lambda > 50 && lambda <= 100) K = 0.95;
//     else if(lambda > 35 && lambda <= 50 ) K = 0.90;
//     else if(lambda > 20 && lambda <= 35 ) K = 0.85;
//     else if(lambda > 10 && lambda <= 20 ) K = 0.75;
//     else if(lambda >  5 && lambda <= 10 ) K = 0.65;
//     else                                  K = 0.60;
//     return K;
// }
//
// double SNiP2_01_07_Cxema14_Ce_x_bez(double deltaD,double Re)
// {
//     if (deltaD >=0.05)
//         return 1.2;
//     if (Re < 4*100000)
//         return 1.2;
//     if (Re > 32 * 100000)
//         return 1.2;
//     double a_Cx[9][4]=
//     {
//     {1.2, 1.2, 1.2, 1.2},//1.5
//     {0.4, 0.4, 0.6, 1.2},//4
//     {0.4, 0.5, 0.7, 1.2},//8
//     {0.5, 0.6, 0.8, 1.2},//12
//     {0.6, 0.7, 0.9, 1.2},//16
//     {0.6, 0.7, 0.9, 1.2},//20
//     {0.6, 0.7, 0.9, 1.2},//24
//     {0.7, 0.8, 1.0, 1.2},//28
//     {0.7, 0.8, 1.0, 1.2} //32
//     };
//     double a_Re[9] = {1.5,4,8,12,16,20,24,28,32};
//     type_LLU i = 0,j =0;
//     for(i=1;i<9;i++)
//         if(a_Re[i]* 100000 >= Re && Re > a_Re[i-1]* 100000) j = i;
//     if (deltaD < 0.0001)
//         return a_Cx[j][0];
//     if (0.001 >= deltaD && deltaD > 0.0001)
//         return a_Cx[j][1];
//     if (0.01 >= deltaD && deltaD > 0.001)
//         return a_Cx[j][2];
//     return a_Cx[j][2]+(a_Cx[j][3]-a_Cx[j][2])*(deltaD - 0.01)/(0.05-0.01);
// }
//
// double SNiP2_01_07_Cxema2_Ce3( double h1, double l, double b, bool OUT=false)
// {
//     double Ce3 = 1e10;
//     if( b/l < 1 )
//     {
//         if     ( h1/l <= 0.5 ) Ce3 = -0.4;
//         else if( h1/l >  2.0 ) Ce3 = -0.6;
//         else                   Ce3 = -0.5;
//     }
//     else
//     {
//         if     ( h1/l <= 0.5 ) Ce3 = -0.5;
//         else if( h1/l >  2.0 ) Ce3 = -0.6;
//         else                   Ce3 = -0.6;
//     }
//     return Ce3;
// }
//
//
//
// enum _Struhale{_Struhale_Cylinder,_Struhale_Rectangle};
// void Printf(_Struhale St)
// {
//     switch(St)
//     {
//         case _Struhale_Cylinder:
//             printf("Number of Strahale 0,2\n");
//             break;
//         case _Struhale_Rectangle:
//             printf("Number of Strahale 0,11\n");
//             break;
//         default:
//             print_name("Number of Strahale is not correct");
//             FATAL();
//     }
// }
//
// double SNiP2_01_07_actual_Formula11_11_Vcr(double f, double d, _Struhale St, bool OUT=false)
// {
//     double dSt = 0.0;
//     switch(St)
//     {
//         case _Struhale_Cylinder :  dSt = 0.20; break;
//         case _Struhale_Rectangle:  dSt = 0.11; break;
//         default:
//             print_name("Number of Strahale is not correct");
//             FATAL();
//     }
//     double Vcr = f*d/dSt;
//     if(OUT)
//     {
//         Printf(St);
//         printf("Vcr = f*d/St = %.3f*%.3f/%.2f = %.2f m/s\n",f,d,dSt,Vcr);
//     }
//     return Vcr;
// }
//
// bool SNiP2_01_07_actual_Formula11_12_Check(double Vcr, double Vmax, bool OUT = false)
// {
//     bool check = true;
//     if(Vcr > Vmax) check = false;
//     else check = true;
//     if(OUT)
//     {
//         printf("Vcr > Vmax\n%.2f > %.2f - ",Vcr,Vmax);
//         if(check) printf("Check case with rezonce\n");
//         else      printf("No rezonance\n");
//     }
//     return check;
// }
//
// double SNiP2_01_07_actual_FormulaD_2_1_F_cylinder
//             (double Vcr, double fi, double d, double delta,
//              bool OUT=false)
// {
//     double Cy = 0.3;
//     double Fi = 0.75*CONST_M_PI*pow(Vcr,2.)*Cy*fi*d/delta;
//     if(OUT)
//     {
//         printf("Fi = 0.75*PI*Vcr^2*Cy*fi*d/delta=0.75*%.3f*%.3f^2*%.3f*%.3f*%.3f/%.3f=%.3fN/m\n",
//                CONST_M_PI,Vcr,Cy,fi,d,delta,Fi);
//     }
//     return Fi;
// }
//
//
// double SNiP2_01_07_Table_19(double L, bool OUT = true)
// {
//     double fu = 0;
//          if(L <=  1.0) fu = 1./120.;
//     else if(L <=  3.0) fu = 1./150.;
//     else if(L <=  6.0) fu = 1./200.;
//     else if(L <= 12.0) fu = 1./250.;
//     else               fu = 1./300.;
//     if(OUT)
//     {
//         printf("Maximum deflection for L = %.1f m is %.2f mm(fu= %.5f)\n",L,L*fu*1e3,fu);
//     }
//     return fu;
// }
//

//
//    type_LLU i;
//    for( i=0; i < (vn-1) ; i++ )
//    {
//        printf("\nРасчет уровня номер %u .\n", i+1 );
//        printf("\nРасчет производится по схеме 12б приложения 4 [__].\n");
//        double hi = H[i+1];
//        printf("Значение h1/d = %.3f.\n",H[vn-1]/d);
//        printf("Определение числа Рейнольдса(Re)\n");
//        double kz  = Veter_BaseFunction::Get_K(hi,Z);
//        printf("Значение коэффициента K для определения Re на высоте %.3f м составляет %.3f.\n",hi,kz);
//        printf("Расчет производится по схеме 12а приложения 4 [__].\n");
//        double _Re = Veter_BaseFunction::Get_Re(d,Wo,kz,true);
//        if ( _Re > 4e5) printf("Re = %.3f  > 400000 - Условие выполнено.\n",_Re);
//        else { FATAL(); printf("Re <= 400000.\n\n\n\n"); return; }
//        double K  = Veter_BaseFunction::Get_K(hi,Z);//k(H1,H2,Z);///ДОПУСКАЕТЬСЯ УМЕНЬШИТЬ ИСПОЛЬЗУЯ АППРОКСИМАЦИЮ
//        type_LLU j;
//        for(j=0;j<((type_LLU)(gn/2+0.6)); j++)
//        {
//            printf("\nУчасток №%u\n",j);
//            double angle1 = j*(360./gn);
//            double angle2 = (j+1)*(360./gn);
//            if(middle == false)
//            {
//                angle1 -= (360./gn)/2.;
//                angle2 -= (360./gn)/2.;
//            }
//            printf("Начальный угол - %.2f град.\n",angle1);
//            printf("Конечный  угол - %.2f град.\n",angle2);
//            double Ce1 = Veter_BaseFunction::Get_Ce1_Cxema12b(angle2,angle1,H[vn-1],d);
//            printf("Определение Wn:\n");
//            Veter_BaseFunction::Wn(Wo,Ce1,K,true);
//        }
//        if(middle == false)
//        {
//            printf("\nУчасток №%u\n",j);
//            double angle1 = j*(360./gn);
//            double angle2 = (j+1)*(360./gn);
//            if(middle == false)
//            {
//                angle1 -= (360./gn)/2.;
//                angle2 -= (360./gn)/2.;
//            }
//            printf("Начальный угол - %.2f град.\n",angle1);
//            printf("Конечный  угол - %.2f град.\n",angle2);
//            double Ce1 = Veter_BaseFunction::Get_Ce1_Cxema12b(angle2,angle1,H[vn-1],d);
//            printf("Определение Wn:\n");
//            Veter_BaseFunction::Wn(Wo,Ce1,K,true);
//        }
//    }

// void WindLoad::Add_Height()
// {
//     height->Sort(func);
//     double H_min = height->Get(0);
//     double H_max = height->Get(height->GetSize()-1);
//     double  Z[14] ={
//         0  ,5  ,10 ,20 ,40 , 60,
//         80 ,100,150,200,250,300,
//         350,480};
//     for(type_LLU i=0;i<14;i++)
//     {
//         if(H_min < Z[i] && Z[i] <H_max)
//             height->Add(Z[i]);
//     }
//     height->Sort(func);
// }
//
// void WindLoad::Printf_K_Dzeta()
// {
//     ::Printf(zone);
//     if(height->GetSize() == 0)return;
//     printf("Height,m\tk\tDzeta\n");
//     for(type_LLU i = 0; i<height->GetSize();i++)
//     {
//         printf("%3.3f\t",height->Get(i));
//         printf("%.3f\t",SNiP2_01_07_table6_K    (height->Get(i),zone));
//         printf("%.3f"  ,SNiP2_01_07_table7_Dzeta(height->Get(i),zone));
//         printf("\n");
//     }
// }
//
//
//
// void WindLoad::CalculateCyl
//                    (double Diameter, type_LLU NumberOfCutting, Array<double> *H, Array<double> *Hz,
//                     double _Wo,
//                     Wind_Log_Decriment LD,
//                     Wind_Zone Zone)
// {
//     height    = H;
//     Wo        = _Wo;
//     zone      = Zone;
//     Log_Decriment = LD;
//     frequency = Hz;
//     //Add_Height();//You can add
//
//     printf("Diameter: %.3f m\n", Diameter);
//     Printf_K_Dzeta();
//
//     Printf(Log_Decriment);
//     double summ_ksi = 0;
//     printf("Natural frecuency\tEta\tKsi\n");
//     for(type_LLU i = 0;i<frequency->GetSize();i++)
//     {
//         double f   = frequency->Get(i);
//         double eta = SNiP2_01_07_p6_7b_Eta(Wo,f);
//         double ksi = SNiP2_01_07_Pics2_Ksi(Log_Decriment,eta);
//         summ_ksi   = sqrt(pow(ksi,2)+pow(summ_ksi,2));
//         printf("%.3fHz\t%.3f\t%.3f\n",f,eta,ksi);
//     }
//     printf("Ksi for calculation: %.3f\n", summ_ksi);
//
//     height->Sort(func);
//     double H_max = height->Get(height->GetSize()-1);
//     double Re    = SNiP2_01_07_Schema12a_Re(Diameter,Wo,
//                    SNiP2_01_07_table6_K    (H_max,zone), true);
//
//     printf("Calculation: average component Wm, N/m2\n");
//     printf("alpha\tCe\tCoeff.K\n");
//     printf("\t\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",SNiP2_01_07_table6_K(height->Get(j),zone));
//     printf("\n");
//     Array<double> Angle;Angle.SetSize(NumberOfCutting+1);
//     Array<double> Wm;Wm.SetSize(height->GetSize()*Angle.GetSize());
//     for(type_LLU i=0;i<Angle.GetSize();i++)
//     {
//         Angle.Set(i,i*180/(Angle.GetSize()-1));
//         double Ce = SNiP2_01_07_Schema12b_Ce1(Angle.Get(i), H_max, Diameter, false);
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wm.Set(i*height->GetSize()+j,SNiP2_01_07_Formula6_Wn(Wo,Ce,SNiP2_01_07_table6_K(height->Get(j),zone)));
//     }
//     for(type_LLU i=0;i<Angle.GetSize();i++)
//     {
//         printf("%3.1f\t",Angle.Get(i));
//         double Ce = SNiP2_01_07_Schema12b_Ce1(Angle.Get(i), H_max, Diameter, false);
//         printf("%1.2f\t",Ce);
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wm.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: pulsing component Wp, N/m2\n");
//     double ro = Diameter;
//     double hi = H_max;
//     double Eps = SNiP2_01_07_Table9_Epsilon(ro, hi);
//     printf("from table 9 SNiP 2.01.07\n");
//     printf("ro = %.3fm\tho = %.3fm\tEps = %.2f\n",ro,hi,Eps);
//
//     Array<double> Wp;Wp.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Angle.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wp.Set(i*height->GetSize()+j,SNiP2_01_07_Formula9_Wp( Wm.Get(i*height->GetSize()+j), SNiP2_01_07_table7_Dzeta(H_max,zone),summ_ksi, Eps));
//     }
//     printf("alpha\tCe\tCoeff.Dzeta\n");
//     printf("\t\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",SNiP2_01_07_table7_Dzeta(height->Get(j),zone));
//     printf("\n");
//     for(type_LLU i=0;i<Angle.GetSize();i++)
//     {
//         printf("%3.1f\t",Angle.Get(i));
//         double Ce = SNiP2_01_07_Schema12b_Ce1(Angle.Get(i), H_max, Diameter, false);
//         printf("%1.2f\t",Ce);
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wp.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: summary component Wsum, N/m2\n");
//     Array<double> Wsum;Wsum.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Angle.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wsum.Set(i*height->GetSize()+j,Wm.Get(i*height->GetSize()+j)+Wp.Get(i*height->GetSize()+j));
//     }
//     printf("alpha\tCe\tHeight,m\n");
//     printf("\t\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",height->Get(j));
//     printf("\n");
//     for(type_LLU i=0;i<Angle.GetSize();i++)
//     {
//         printf("%3.1f\t",Angle.Get(i));
//         double Ce = SNiP2_01_07_Schema12b_Ce1(Angle.Get(i), H_max, Diameter, false);
//         printf("%1.2f\t",Ce);
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wsum.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: summary force\n");
//     Array<double> Fx,Fy; Fx.SetSize(Angle.GetSize()-1); Fy.SetSize(Angle.GetSize()-1);
//     printf("Angle\tFaw, N\tFx, N\tFy, N\n");
//     for(type_LLU i=0;i<Angle.GetSize()-1;i++)
//     {
//         printf("%03.1f - %3.1f\t",Angle.Get(i),Angle.Get(i+1));
//         double Faw = 0;
//         for(type_LLU j = 1;j<height->GetSize();j++)
//         {
//             double L = CONST_M_PI*Diameter*fabs(Angle.Get(i)-Angle.Get(i+1))/360;
//             double Waw =
//                   (Wsum.Get((i+0)*height->GetSize()+j-0)+
//                    Wsum.Get((i+1)*height->GetSize()+j-0)+
//                    Wsum.Get((i+0)*height->GetSize()+j-1)+
//                    Wsum.Get((i+1)*height->GetSize()+j-1))/4;
//             Faw += Waw * (height->Get(j)-height->Get(j-1))*L;
//         }
//         printf("%.1f\t",Faw);
//         double AlphaAw = (Angle.Get(i)+Angle.Get(i+1))/2.;
//         Fx.Set(i,Faw*cos(AlphaAw*CONST_M_PI/180));
//         Fy.Set(i,Faw*sin(AlphaAw*CONST_M_PI/180));
//         printf("%.1f\t",Fx.Get(i));
//         printf("%.1f\t",Fy.Get(i));
//         printf("\n");
//     }
//     double summ_Fx=0,summ_Fy=0;
//     for(type_LLU i=0;i<Fx.GetSize();i++)
//     {
//         summ_Fx += Fx.Get(i);
//         summ_Fy += Fy.Get(i);
//     }
//     printf("Summary of Fx: %5.1fN\n",summ_Fx);
//     printf("Summary of Fy: %5.1fN\n",summ_Fy);
//
//     // Rezonce check
//     printf("\nRezonce check\n");
//     double Vmax = SNiP2_01_07_actual_Formula11_13_Vmax(Wo, SNiP2_01_07_table6_K(H_max,zone),true);
//     bool Check_Rezonance = false;
//     for(type_LLU i =0;i<frequency->GetSize();i++)
//     {
//         double Vcr = SNiP2_01_07_actual_Formula11_11_Vcr(frequency->Get(i),Diameter,Wind_Struhale_Cylinder,true);
//         if(SNiP2_01_07_actual_Formula11_12_Check(Vcr, Vmax, true)) Check_Rezonance = true;
//     }
//     if(Check_Rezonance) printf("Check rezonance\n");
//     else printf("No rezonance in case\n");
//
//     ////
//     print_name("\n\nCalculate wind load(schema 14 Appendix 4 SNiP 2.01.07)\n");
//     double l = H_max;
//     double b = Diameter;
//     double lambda = l/b;
//     double lambdae = 2*lambda;
//     double K = SNiP2_01_07_Cxema13_table1_K(lambdae);
//     double Ce_bez = SNiP2_01_07_Cxema14_Ce_x_bez(0.001,Re);
//     printf("see schema4 Appendix 4:\n");
//     printf("l = %.3fm\tb = %.3fm\tlambdae = %.2f\tK = %.2f\nRe = %.1fe5\tCe_bez = %.1f\tsumm_ksi = %.2f\n",l,b,lambdae,K,Re*1e-5,Ce_bez,summ_ksi);
//     Wm.SetSize(height->GetSize());
//     for(type_LLU j = 0;j<height->GetSize();j++)
//         Wm.Set(j,SNiP2_01_07_Formula6_Wn(Wo,Ce_bez,SNiP2_01_07_table6_K(height->Get(j),zone)));
//     Wp.SetSize(Wm.GetSize());
//     for(type_LLU j = 0;j<height->GetSize();j++)
//         Wp.Set(j,SNiP2_01_07_Formula9_Wp( Wm.Get(j), SNiP2_01_07_table7_Dzeta(H_max,zone),summ_ksi, Eps));
//     Wsum.SetSize(Wm.GetSize());
//     for(type_LLU j = 0;j<height->GetSize();j++)
//         Wsum.Set(j,Wm.Get(j)+Wp.Get(j));
//     Array<double>Qsum;Qsum.SetSize(Wm.GetSize());
//     for(type_LLU j = 0;j<height->GetSize();j++)
//         Qsum.Set(j,Wsum.Get(j)*Diameter);
//     printf("Height,m\tWm,N/m2\tK\tWp,N/m2\tWsum,N/m2\tQsum, N/m\n");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//     {
//         printf("%.3f\t%.1f\t%.1f\t%.2f\t%.1f\t%.1f\n",height->Get(j),Wm.Get(j),SNiP2_01_07_table6_K(height->Get(j),zone),Wp.Get(j),Wsum.Get(j),Qsum.Get(j));
//     }
//
//     printf("\n");
//     printf("Calculation: summary force\n");
//     summ_Fx = 0;
//     for(type_LLU j = 1;j<height->GetSize();j++)
//     {
//         summ_Fx += (height->Get(j)-height->Get(j-1))*(Wsum.Get(j)+Wsum.Get(j-1))/2.*Diameter;
//     }
//     printf("Summary of Fx: %5.1fN\n",summ_Fx);
// };
//
// void WindLoad::CalculateCyl
//                 ( MSH &mesh, _LOAD * load,double HeightMin,double HeightMax,double PositionX0,double PositionZ0,double precition,
//                 double Diameter, type_LLU NumberOfCutting, Array<double> *Hz,
//                 double _Wo,
//                 Wind_Log_Decriment LD,
//                 Wind_Zone Zone,
//                 UNIT_FORCE  uf,
//                 UNIT_LENGHT ul)
// {
//     int InvertForcesOnCylinder = +1.00;
//     if(_Wo < 0){_Wo = -_Wo;InvertForcesOnCylinder = -1.00;}
//     //height    = H;
//     Wo        = _Wo;
//     zone      = Zone;
//     Log_Decriment = LD;
//     frequency = Hz;
//     double UnitFactor = 1;
//     switch(uf)
//     {
//         case(UNIT_FORCE_N ): UnitFactor *= 1;    break;
//         case(UNIT_FORCE_KN): UnitFactor *= 0.001;break;
//         default: print_name("FATAL ERROR in UNIT_FORCE");FATAL();
//     };
//     switch(ul)
//     {
//         case(UNIT_LENGHT_METER ): UnitFactor *= 1;break;
//         case(UNIT_LENGHT_MMS   ): UnitFactor *= 1e-6;break;
//         default: print_name("FATAL ERROR in UNIT_METER");FATAL();
//     }
//
//     /// Find cylinder
//     Array<type_LLU> num_nodes  ;
//     Array<type_LLU> num_element;
//     mesh.FindCylinder(Diameter,HeightMin,HeightMax,PositionX0,PositionZ0,precition,num_nodes,num_element);
//
//     /// Find plates on cylinder
//     Array<type_LLU> num_element_cylinder;
//     for(type_LLU i=0;i<num_element.GetSize();i++)
//     {
//         Node n1,n2,n0;
//         for(type_LLU j=0;j<num_nodes.GetSize();j++)
//             if(mesh.elements.Get(num_element.Get(i)).node[0] == mesh.nodes.Get(num_nodes.Get(j)).Number)
//             { n0 =  mesh.nodes.Get(num_nodes.Get(j)); break; }
//         for(type_LLU j=0;j<num_nodes.GetSize();j++)
//             if(mesh.elements.Get(num_element.Get(i)).node[1] == mesh.nodes.Get(num_nodes.Get(j)).Number)
//             { n1 =  mesh.nodes.Get(num_nodes.Get(j)); break; }
//         for(type_LLU j=0;j<num_nodes.GetSize();j++)
//             if(mesh.elements.Get(num_element.Get(i)).node[2] == mesh.nodes.Get(num_nodes.Get(j)).Number)
//             { n2 =  mesh.nodes.Get(num_nodes.Get(j)); break; }
//         if(GET_PLATE_POSITION(n0,n1,n2) != PLANE_POSITION_Y && (mesh.elements.Get(num_element.Get(i)).ElmType == ELEMENT_TYPE_QUADRANGLE || mesh.elements.Get(num_element.Get(i)).ElmType == ELEMENT_TYPE_TRIANGLE))
//             num_element_cylinder.Add(num_element.Get(i));
//     }
//
//     /// Draw resilts
//
//     /// Calculate wind load
//     /// _LOAD - load[4]
//     //  load[0] - Wind +X
//     //  load[1] - Wind -X
//     //  load[2] - Wind +Z
//     //  load[3] - Wind -Z
//     for(type_LLU WL = 0; WL < 4 ;WL++)
//     {
//         double Fx = 0;
//         double Fy = 0;
//         double AreaAllPates = 0;
//         for(type_LLU i=0;i<num_element_cylinder.GetSize();i++)
//         {
//             // Find coordination nodes
//             Node n[4];
//             type_LLU nmax=0;
//             switch(mesh.elements.Get(num_element_cylinder.Get(i)).ElmType)
//             {
//                 case ELEMENT_TYPE_QUADRANGLE: nmax = 4; break;
//                 case ELEMENT_TYPE_TRIANGLE  : nmax = 3; break;
//                 default: print_name("WARNING in CalculateCyl"); Printf(mesh.elements.Get(num_element_cylinder.Get(i)).ElmType); FATAL();
//             }
//             for(type_LLU k=0;k<nmax;k++)
//                 for(type_LLU j=0;j<num_nodes.GetSize();j++)
//                     if(mesh.elements.Get(num_element_cylinder.Get(i)).node[k] == mesh.nodes.Get(num_nodes.Get(j)).Number)
//                     {
//                         n[k] =  mesh.nodes.Get(num_nodes.Get(j));
//                         break;
//                     }
//             // Find center point
//             double Xc = 0;
//             double Yc = 0;
//             double Zc = 0;
//             for(type_LLU k=0;k<nmax;k++)
//             {
//                 Xc += n[k].x;
//                 Yc += n[k].y;
//                 Zc += n[k].z;
//             }
//             Xc /= nmax;
//             Yc /= nmax;
//             Zc /= nmax;
//             // Find angle
//             double angle;
//             switch(WL)
//             {
//                 case 0://load[0] - Wind +X
//                     angle = AngleCoordination(-(Xc-PositionX0),+(Zc-PositionZ0)); break;
//                 case 1://load[0] - Wind -X
//                     angle = AngleCoordination(+(Xc-PositionX0),+(Zc-PositionZ0)); break;
//                 case 2://load[0] - Wind +Z
//                     angle = AngleCoordination(+(Xc-PositionX0),-(Zc-PositionZ0)); break;
//                 case 3://load[0] - Wind -Z
//                     angle = AngleCoordination(+(Xc-PositionX0),+(Zc-PositionZ0)); break;
//                 default:
//                     print_name("Warning in switch(WL):CalculateCyl");
//                     FATAL();
//             }
//             angle = GRADIANS(angle);
//             // Find height
//             // Find wind load
//             double summ_ksi = 0;
//             if(frequency->GetSize() > 0)
//             {
//                 for(type_LLU i = 0;i<frequency->GetSize();i++)
//                 {
//                     double f   = frequency->Get(i);
//                     double eta = SNiP2_01_07_p6_7b_Eta(Wo,f);
//                     double ksi = SNiP2_01_07_Pics2_Ksi(Log_Decriment,eta);
//                     summ_ksi   = sqrt(pow(ksi,2)+pow(summ_ksi,2));
//                 }
//             }
//             else
//             {
//                 print_name("ERROR in if(frequency->GetSize() > 0)");
//                 FATAL();
//             }
//             double Re    = SNiP2_01_07_Schema12a_Re(Diameter ,Wo,
//                            SNiP2_01_07_table6_K    (HeightMax,zone), false);
//
//             double Wn;
//             double Ce = SNiP2_01_07_Schema12b_Ce1(angle, HeightMax, Diameter, false);
//             Wn = SNiP2_01_07_Formula6_Wn(Wo,Ce,SNiP2_01_07_table6_K(Yc,zone));
//
//             double ro = Diameter;
//             double hi = HeightMax;
//             double Eps = SNiP2_01_07_Table9_Epsilon(ro, hi);
//
//             double Wp;
//             Wp = SNiP2_01_07_Formula9_Wp( Wn, SNiP2_01_07_table7_Dzeta(HeightMax,zone),summ_ksi, Eps);
//
//             double Wsum;
//             Wsum = Wn + Wp;
//
//
//             //Wsum = ROUND(Wsum,6);
//
//             LT.FT = FORCE_TYPE_ELEMENT;
//             LT.FS = FORCE_SYSTEM_LOCAL;
//             LT.FD = FORCE_DIRECTION_X ;
//             LT.NumberOfElement = mesh.elements.Get(num_element_cylinder.Get(i)).Number;
//             LT.value[0] = Wsum*InvertForcesOnCylinder*UnitFactor;
//             load[WL].SLT.Add(LT);
//             // Calculate summary forces
//             double area_element = 0;
//             switch(mesh.elements.Get(num_element_cylinder.Get(i)).ElmType)
//             {
//                 case ELEMENT_TYPE_POINT:
//                 case ELEMENT_TYPE_LINE:
//                 case ELEMENT_TYPE_TETRAHEDRON:
//                     ; break;
//                 case ELEMENT_TYPE_QUADRANGLE: area_element = area_4node(n[0],n[1],n[2],n[3]); break;
//                 case ELEMENT_TYPE_TRIANGLE  : area_element = area_3node(n[0],n[1],n[2]     ); break;
//             }
//             Fx += Wsum*area_element*sin(RADIANS(angle));
//             Fy += Wsum*area_element*cos(RADIANS(angle));
//             AreaAllPates += area_element;
//         }
//         printf("Summary of Fx: %5.1fN\n",Fx);
//         printf("Summary of Fy: %5.1fN\n",Fy);
//         printf("Area of All Plates: %5.3f m2\n\n",AreaAllPates);
//     }
// };
//
// void WindLoad::CalculateConv
//                    (double XX, double YY, Array<double> *H, Array<double> *Hz,
//                     double _Wo,
//                     Wind_Log_Decriment LD,
//                     Wind_Zone Zone)
// {
//     height    = H;
//     Wo        = _Wo;
//     zone      = Zone;
//     Log_Decriment = LD;
//     frequency = Hz;
//     printf("\nWind with convection X\n");
//     CalculateC(XX,YY);
//     printf("\nWind with convection Y\n");
//     CalculateC(YY,XX);
// }
//
// void WindLoad::CalculateC(double XX, double YY)
// {
//     Array<double> Ce;
//     Ce.Add(0.8);
//
//     height->Sort(func);
//     double H_min = height->Get(0);
//     double H_max = height->Get(height->GetSize()-1);
//     Ce.Add(SNiP2_01_07_Cxema2_Ce3( H_max-H_min, XX, YY, true));
//
//     printf("Dimension X: %.3f m\n", XX);
//     printf("Dimension Y: %.3f m\n", YY);
//     Printf_K_Dzeta();
//
//     Printf(Log_Decriment);
//     double summ_ksi = 0;
//     printf("Natural frecuency\tEta\tKsi\n");
//     for(type_LLU i = 0;i<frequency->GetSize();i++)
//     {
//         double f   = frequency->Get(i);
//         double eta = SNiP2_01_07_p6_7b_Eta(Wo,f);
//         double ksi = SNiP2_01_07_Pics2_Ksi(Log_Decriment,eta);
//         summ_ksi   = sqrt(pow(ksi,2)+pow(summ_ksi,2));
//         printf("%.3fHz\t%.3f\t%.3f\n",f,eta,ksi);
//     }
//     printf("Ksi for calculation: %.3f\n", summ_ksi);
//
//     printf("Calculation: average component Wm, N/m2\n");
//     printf("Ce\tCoeff.K\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",SNiP2_01_07_table6_K(height->Get(j),zone));
//     printf("\n");
//     Array<double> Wm;Wm.SetSize(height->GetSize()*Ce.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//         {
//             Wm.Set(i*height->GetSize()+j,SNiP2_01_07_Formula6_Wn(Wo,Ce.Get(i),SNiP2_01_07_table6_K(height->Get(j),zone)));
//             printf("%3.f\t",Wm.Get(i*height->GetSize()+j));
//         }
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: pulsing component Wp, N/m2\n");
//     double ro = YY;
//     double hi = H_max;
//     double Eps = SNiP2_01_07_Table9_Epsilon(ro, hi);
//     printf("from table 9 SNiP 2.01.07\n");
//     printf("ro = %.3fm\tho = %.3fm\tEps = %.2f\n",ro,hi,Eps);
//     Array<double> Wp;Wp.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wp.Set(i*height->GetSize()+j,SNiP2_01_07_Formula9_Wp( Wm.Get(i*height->GetSize()+j), SNiP2_01_07_table7_Dzeta(H_max,zone),summ_ksi, Eps));
//     }
//     printf("Ce\tCoeff.Dzeta\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",SNiP2_01_07_table7_Dzeta(height->Get(j),zone));
//     printf("\n");
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%1.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wp.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: summary component Wsum, N/m2\n");
//     Array<double> Wsum;Wsum.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wsum.Set(i*height->GetSize()+j,Wm.Get(i*height->GetSize()+j)+Wp.Get(i*height->GetSize()+j));
//     }
//     printf("Ce\tHeight,m\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",height->Get(j));
//     printf("\n");
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%1.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wsum.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: summary force\n");
//     double Fx=0,Fy=0;
//     for(type_LLU j = 1;j<height->GetSize();j++)
//     {
//         Fx += (Wsum.Get(0*height->GetSize()+j-0)+
//                Wsum.Get(0*height->GetSize()+j-1)-
//                Wsum.Get(1*height->GetSize()+j-0)-
//                Wsum.Get(1*height->GetSize()+j-1))/2.*
//               (height->Get(j)-height->Get(j-1))*XX;
//         Fy += (Wsum.Get(1*height->GetSize()+j)+
//                Wsum.Get(1*height->GetSize()+j-1))/2.*
//               (height->Get(j)-height->Get(j-1))*YY;
//     }
//     printf("Summary of Fx: %5.1fN\n",Fx);
//     printf("Summary of Fy: %5.1fN\n",Fy);
//
//     // Rezonce check
//     printf("\nRezonce check\n");
//     double Vmax = SNiP2_01_07_actual_Formula11_13_Vmax(Wo, SNiP2_01_07_table6_K(H_max,zone),true);
//     bool Check_Rezonance = false;
//     for(type_LLU i =0;i<frequency->GetSize();i++)
//     {
//         double Vcr = SNiP2_01_07_actual_Formula11_11_Vcr(frequency->Get(i),YY,Wind_Struhale_Rectangle,true);
//         if(SNiP2_01_07_actual_Formula11_12_Check(Vcr, Vmax, true)) Check_Rezonance = true;
//     }
//     if(Check_Rezonance) printf("Check rezonance\n");
//     else printf("No rezonance in case\n");
// }
//
// void WindLoad::CalculateConv
//                 (MSH &mesh, ,
//                 Node _n1, Node _n2, double precition,
//                 Array<double> *Hz,
//                 double _Wo,
//                 Wind_Log_Decriment LD,
//                 Wind_Zone Zone,
//                 UNIT_FORCE  uf,
//                 UNIT_LENGHT ul)
// {
//     double UnitFactor = 1;
//     switch(uf)
//     {
//         case(UNIT_FORCE_N ): UnitFactor *= 1;break;
//         case(UNIT_FORCE_KN): UnitFactor *= 0.001;break;
//         default: print_name("FATAL ERROR in UNIT_FORCE");FATAL();
//     };
//     switch(ul)
//     {
//         case(UNIT_LENGHT_METER ): UnitFactor *= 1;break;
//         case(UNIT_LENGHT_MMS   ): UnitFactor *= 1e-6;break;
//         default: print_name("FATAL ERROR in UNIT_FORCE");FATAL();
//     }
//     Node n1,n2;
//     n1.x = min(_n1.x,_n2.x);n2.x = max(_n1.x,_n2.x);
//     n1.y = min(_n1.y,_n2.y);n2.y = max(_n1.y,_n2.y);
//     n1.z = min(_n1.z,_n2.z);n2.z = max(_n1.z,_n2.z);
//     Wo        = _Wo;
//     zone      = Zone;
//     Log_Decriment = LD;
//     frequency = Hz;
//
//     type_LLU j;
//
//     for(type_LLU k=0;k<4;k++)
//     {
//         double XX,YY;
//         if(k == 0 || k == 1) {XX = fabs(n2.z-n1.z); YY = fabs(n2.x-n1.x);}
//         else                 {XX = fabs(n2.x-n1.x); YY = fabs(n2.z-n1.z);}
//         Array<double> Ce;
//         Ce.Add(0.8);
//
//         //height->SetSize(0);
//         double H_min = min(n1.y,n2.y);
//         double H_max = max(n1.y,n2.y);
//         height = new Array<double>;
//         height->SetSize(5+1);
//         for(type_LLU u=0;u<height->GetSize();u++)
//             height->Set(u,H_min+(H_max-H_min)/(height->GetSize()-1)*u);
//             printf("%u\t%.9f m\n",u,height->Get(u));
//         Ce.Add(SNiP2_01_07_Cxema2_Ce3( H_max-H_min, XX, YY, DEBUG));
//
//         double summ_ksi = 0;
//         for(type_LLU i = 0;i<frequency->GetSize();i++)
//         {
//             double f   = frequency->Get(i);
//             double eta = SNiP2_01_07_p6_7b_Eta(Wo,f);
//             double ksi = SNiP2_01_07_Pics2_Ksi(Log_Decriment,eta);
//             summ_ksi   = sqrt(pow(ksi,2)+pow(summ_ksi,2));
//         }
//         Array<double> Wm;Wm.SetSize(height->GetSize()*Ce.GetSize());
//         for(type_LLU i=0;i<Ce.GetSize();i++)
//         {
//             for(type_LLU j = 0;j<height->GetSize();j++)
//             {
//                 Wm.Set(i*height->GetSize()+j,SNiP2_01_07_Formula6_Wn(Wo,Ce.Get(i),SNiP2_01_07_table6_K(height->Get(j),zone)));
//             }
//         }
//
//         double ro = YY;
//         double hi = H_max;
//         double Eps = SNiP2_01_07_Table9_Epsilon(ro, hi);
//         Array<double> Wp;Wp.SetSize(Wm.GetSize());
//         for(type_LLU i=0;i<Ce.GetSize();i++)
//         {
//             for(type_LLU j = 0;j<height->GetSize();j++)
//                 Wp.Set(i*height->GetSize()+j,SNiP2_01_07_Formula9_Wp( Wm.Get(i*height->GetSize()+j), SNiP2_01_07_table7_Dzeta(H_max,zone),summ_ksi, Eps));
//         }
//         Array<double> Wsum;Wsum.SetSize(Wm.GetSize());
//         for(type_LLU i=0;i<Ce.GetSize();i++)
//         {
//             for(type_LLU j = 0;j<height->GetSize();j++)
//                 Wsum.Set(i*height->GetSize()+j,Wm.Get(i*height->GetSize()+j)+Wp.Get(i*height->GetSize()+j));
//         }
//
//         Node nn1,nn2;
//         LT;
//         Array<type_LLU> num_nodes  ;
//         Array<type_LLU> num_element;
//
//         // Add to
//         LT.FT = FORCE_TYPE_ELEMENT;
//         LT.FS = FORCE_SYSTEM_GLOBAL;
//         LT.FD = FORCE_DIRECTION_X ;
//         // PLATES 1
//         nn1.x  = n1.x; nn2.x = n1.x;
//         nn1.y  = n1.y; nn2.y = n2.y;
//         nn1.z  = n1.z; nn2.z = n2.z;
//         num_nodes  .SetSize(0);
//         num_element.SetSize(0);
//         mesh.FindPlates(nn1,nn2,PLANE_POSITION_X,precition,num_nodes,num_element);
//
//         type_LLU uh;
//         for(uh=0;uh<num_element.GetSize();uh++)
//         {
//             Node           np[3];
//             type_LLU Number_n[3];
//             for(j=0;j<3;j++)
//                 Number_n[j] = mesh.elements.Get(num_element.Get(uh)).node[j];
//             for(j=0;j<mesh.nodes.GetSize();j++)
//             {
//                 if(mesh.nodes.Get(j).Number == Number_n[0]) np[0] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[1]) np[1] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[2]) np[2] = mesh.nodes.Get(j);
//             }
//             double Ycenter = (np[0].y+np[1].y+np[2].y)/3.;
//             for(j=0;j<height->GetSize();j++)
//                 if(height->Get(j)>=Ycenter)
//                 {break;}
//             LT.NumberOfElement = mesh.elements.Get(num_element.Get(uh)).Number;
//             switch(k)
//             {
//                 case(0): LT.value[0] = +Wsum.Get(0*height->GetSize()+j)*UnitFactor; break;
//                 case(1): LT.value[0] = +Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(2): LT.value[0] = +Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(3): LT.value[0] = +Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//             }
//             //LT.value[0] = ROUND(LT.value[0],6);
//             load[k].SLT.Add(LT);
//         }
//         // PLATES 2
//         nn1.x  = n2.x; nn2.x = n2.x;
//         nn1.y  = n1.y; nn2.y = n2.y;
//         nn1.z  = n1.z; nn2.z = n2.z;
//         num_nodes  .SetSize(0);
//         num_element.SetSize(0);
//         mesh.FindPlates(nn1,nn2,PLANE_POSITION_X,precition,num_nodes,num_element);
//
//         for(uh=0;uh<num_element.GetSize();uh++)
//         {
//             Node           np[3];
//             type_LLU Number_n[3];
//             for(j=0;j<3;j++)
//                 Number_n[j] = mesh.elements.Get(num_element.Get(uh)).node[j];
//             for(j=0;j<mesh.nodes.GetSize();j++)
//             {
//                 if(mesh.nodes.Get(j).Number == Number_n[0]) np[0] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[1]) np[1] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[2]) np[2] = mesh.nodes.Get(j);
//             }
//             double Ycenter = (np[0].y+np[1].y+np[2].y)/3.;
//             for(j=0;j<height->GetSize();j++)
//                 if(height->Get(j)>=Ycenter)
//                 {break;}
//             LT.NumberOfElement = mesh.elements.Get(num_element.Get(uh)).Number;
//             switch(k)
//             {
//                 case(0): LT.value[0] = -Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(1): LT.value[0] = -Wsum.Get(0*height->GetSize()+j)*UnitFactor; break;
//                 case(2): LT.value[0] = -Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(3): LT.value[0] = -Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//             }
//             //LT.value[0] = ROUND(LT.value[0],6);
//             load[k].SLT.Add(LT);
//         }
//
//         // Add to
//         LT.FT = FORCE_TYPE_ELEMENT;
//         LT.FS = FORCE_SYSTEM_GLOBAL;
//         LT.FD = FORCE_DIRECTION_Z ;
//         // PLATES 3
//         nn1.x  = n1.x; nn2.x = n2.x;
//         nn1.y  = n1.y; nn2.y = n2.y;
//         nn1.z  = n1.z; nn2.z = n1.z;
//         num_nodes  .SetSize(0);
//         num_element.SetSize(0);
//         mesh.FindPlates(nn1,nn2,PLANE_POSITION_Z,precition,num_nodes,num_element);
//
//         for(uh=0;uh<num_element.GetSize();uh++)
//         {
//             Node           np[3];
//             type_LLU Number_n[3];
//             for(j=0;j<3;j++)
//                 Number_n[j] = mesh.elements.Get(num_element.Get(uh)).node[j];
//             for(j=0;j<mesh.nodes.GetSize();j++)
//             {
//                 if(mesh.nodes.Get(j).Number == Number_n[0]) np[0] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[1]) np[1] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[2]) np[2] = mesh.nodes.Get(j);
//             }
//             double Ycenter = (np[0].y+np[1].y+np[2].y)/3.;
//             for(j=0;j<height->GetSize();j++)
//                 if(height->Get(j)>=Ycenter)
//                 {break;}
//             LT.NumberOfElement = mesh.elements.Get(num_element.Get(uh)).Number;
//             switch(k)
//             {
//                 case(0): LT.value[0] = +Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(1): LT.value[0] = +Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(2): LT.value[0] = +Wsum.Get(0*height->GetSize()+j)*UnitFactor; break;
//                 case(3): LT.value[0] = +Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//             }
//             //LT.value[0] = ROUND(LT.value[0],6);
//             load[k].SLT.Add(LT);
//         }
//         // PLATES 4
//         nn1.x  = n1.x; nn2.x = n2.x;
//         nn1.y  = n1.y; nn2.y = n2.y;
//         nn1.z  = n2.z; nn2.z = n2.z;
//         num_nodes  .SetSize(0);
//         num_element.SetSize(0);
//         mesh.FindPlates(nn1,nn2,PLANE_POSITION_Z,precition,num_nodes,num_element);
//
//         for(uh=0;uh<num_element.GetSize();uh++)
//         {
//             Node           np[3];
//             type_LLU Number_n[3];
//             for(j=0;j<3;j++)
//                 Number_n[j] = mesh.elements.Get(num_element.Get(uh)).node[j];
//             for(j=0;j<mesh.nodes.GetSize();j++)
//             {
//                 if(mesh.nodes.Get(j).Number == Number_n[0]) np[0] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[1]) np[1] = mesh.nodes.Get(j);
//                 if(mesh.nodes.Get(j).Number == Number_n[2]) np[2] = mesh.nodes.Get(j);
//             }
//             double Ycenter = (np[0].y+np[1].y+np[2].y)/3.;
//             for(j=0;j<height->GetSize();j++)
//                 if(height->Get(j)>=Ycenter)
//                 {break;}
//             LT.NumberOfElement = mesh.elements.Get(num_element.Get(uh)).Number;
//             switch(k)
//             {
//                 case(0): LT.value[0] = -Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(1): LT.value[0] = -Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(2): LT.value[0] = -Wsum.Get(1*height->GetSize()+j)*UnitFactor; break;
//                 case(3): LT.value[0] = -Wsum.Get(0*height->GetSize()+j)*UnitFactor; break;
//             }
//             //LT.value[0] = ROUND(LT.value[0],6);
//             load[k].SLT.Add(LT);
//         }
//     }
// };
//
// void WindLoad::CalculateFrame
//                    (double XX, double YY, Array<double> *H, Array<double> *Hz,
//                     double _Wo,
//                     Wind_Log_Decriment LD,
//                     Wind_Zone Zone)
// {
//     height    = H;
//     Wo        = _Wo;
//     zone      = Zone;
//     Log_Decriment = LD;
//     frequency = Hz;
//     printf("Calculate frame\n");
//     printf("\nWind with frame X\n");
//     CalculateFr(XX,YY);
//     printf("\nWind with frame Y\n");
//     CalculateFr(YY,XX);
// }
//
// void WindLoad::CalculateFr(double XX, double YY)
// {
//     Array<double> Ce;
//     Ce.Add(1.4);
//
//     height->Sort(func);
//     double H_min = height->Get(0);
//     double H_max = height->Get(height->GetSize()-1);
//
//     printf("Dimension X: %.3f m\n", XX);
//     printf("Dimension Y: %.3f m\n", YY);
//     Printf_K_Dzeta();
//
//     Printf(Log_Decriment);
//     double summ_ksi = 0;
//     printf("Natural frecuency\tEta\tKsi\n");
//     for(type_LLU i = 0;i<frequency->GetSize();i++)
//     {
//         double f   = frequency->Get(i);
//         double eta = SNiP2_01_07_p6_7b_Eta(Wo,f);
//         double ksi = SNiP2_01_07_Pics2_Ksi(Log_Decriment,eta);
//         summ_ksi   = sqrt(pow(ksi,2)+pow(summ_ksi,2));
//         printf("%.3fHz\t%.3f\t%.3f\n",f,eta,ksi);
//     }
//     printf("Ksi for calculation: %.3f\n", summ_ksi);
//
//     printf("Calculation: average component Wm, N/m2\n");
//     printf("Ce\tCoeff.K\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",SNiP2_01_07_table6_K(height->Get(j),zone));
//     printf("\n");
//     Array<double> Wm;Wm.SetSize(height->GetSize()*Ce.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//         {
//             Wm.Set(i*height->GetSize()+j,SNiP2_01_07_Formula6_Wn(Wo,Ce.Get(i),SNiP2_01_07_table6_K(height->Get(j),zone)));
//             printf("%3.f\t",Wm.Get(i*height->GetSize()+j));
//         }
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: pulsing component Wp, N/m2\n");
//     double ro = YY;
//     double hi = H_max;
//     double Eps = SNiP2_01_07_Table9_Epsilon(ro, hi);
//     printf("from table 9 SNiP 2.01.07\n");
//     printf("ro = %.3fm\tho = %.3fm\tEps = %.2f\n",ro,hi,Eps);
//     Array<double> Wp;Wp.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wp.Set(i*height->GetSize()+j,SNiP2_01_07_Formula9_Wp( Wm.Get(i*height->GetSize()+j), SNiP2_01_07_table7_Dzeta(H_max,zone),summ_ksi, Eps));
//     }
//     printf("Ce\tCoeff.Dzeta\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",SNiP2_01_07_table7_Dzeta(height->Get(j),zone));
//     printf("\n");
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%1.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wp.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: summary component Wsum, N/m2\n");
//     Array<double> Wsum;Wsum.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wsum.Set(i*height->GetSize()+j,Wm.Get(i*height->GetSize()+j)+Wp.Get(i*height->GetSize()+j));
//     }
//     printf("Ce\tHeight,m\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",height->Get(j));
//     printf("\n");
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%1.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.f\t",Wsum.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     printf("\n");
//     printf("Calculation: summary component Qsum,N/m\n");
//     Array<double> Qsum;Qsum.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Qsum.Set(i*height->GetSize()+j,Wm.Get(i*height->GetSize()+j)*XX);
//     }
//     printf("Ce\tHeight,m\n");
//     printf("\t");
//     for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%+1.3f\t",height->Get(j));
//     printf("\n");
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         printf("%1.2f\t",Ce.Get(i));
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             printf("%3.1f\t",Qsum.Get(i*height->GetSize()+j));
//         printf("\n");
//     }
//
//     // Rezonce check
//     printf("\nRezonce check\n");
//     double Vmax = SNiP2_01_07_actual_Formula11_13_Vmax(Wo, SNiP2_01_07_table6_K(H_max,zone),true);
//     bool Check_Rezonance = false;
//     for(type_LLU i =0;i<frequency->GetSize();i++)
//     {
//         double Vcr = SNiP2_01_07_actual_Formula11_11_Vcr(frequency->Get(i),YY,Wind_Struhale_Rectangle,true);
//         if(SNiP2_01_07_actual_Formula11_12_Check(Vcr, Vmax, true)) Check_Rezonance = true;
//     }
//     if(Check_Rezonance) printf("Check rezonance\n");
//     else printf("No rezonance in case\n");
// }
//
//
//
//
// void WindLoad::CalculateFrame
//                 (MSH &mesh, *load, Array<MEMBER_PROPERTY> smp,
//                 Array<double> *Hz,
//                 double _Wo,
//                 Wind_Log_Decriment LD,
//                 Wind_Zone Zone,
//                 UNIT_FORCE  uf,
//                 UNIT_LENGHT ul)
// {
//
//     double UnitFactor = 1;
//     switch(uf)
//     {
//         case(UNIT_FORCE_N ): UnitFactor *= 1;break;
//         case(UNIT_FORCE_KN): UnitFactor *= 0.001;break;
//         default: print_name("FATAL ERROR in UNIT_FORCE");FATAL();
//     };
//     switch(ul)
//     {
//         case(UNIT_LENGHT_METER ): UnitFactor *= 1;break;
//         case(UNIT_LENGHT_MMS   ): UnitFactor *= 1e-6;break;
//         default: print_name("FATAL ERROR in UNIT_FORCE");FATAL();
//     }
//     Wo        = _Wo;
//     zone      = Zone;
//     Log_Decriment = LD;
//     frequency = Hz;
//
//     // HEIGHT //
//     double H_min = 20;
//     double H_max = 0;
//     for(type_LLU i=0;i<mesh.nodes.GetSize();i++)
//     {
//         if(H_min > mesh.nodes.Get(i).y) H_min = mesh.nodes.Get(i).y;
//         if(H_max < mesh.nodes.Get(i).y) H_max = mesh.nodes.Get(i).y;
//     }
//     height = new Array<double>;
//     height->SetSize(5+1);
//     for(type_LLU u=0;u<height->GetSize();u++)
//         height->Set(u,H_min+(H_max-H_min)/(height->GetSize()-1)*u);
//
//     // SIZE PROFILE //
//     // 300 mm       //
//     double XX, YY=XX = 0.300;
//
//
//     Array<double> Ce;
//     Ce.Add(1.4);
//
//     height->Sort(func);
//     H_min = height->Get(0);
//     H_max = height->Get(height->GetSize()-1);
//
//     double summ_ksi = 0;
//     for(type_LLU i = 0;i<frequency->GetSize();i++)
//     {
//         double f   = frequency->Get(i);
//         double eta = SNiP2_01_07_p6_7b_Eta(Wo,f);
//         double ksi = SNiP2_01_07_Pics2_Ksi(Log_Decriment,eta);
//         summ_ksi   = sqrt(pow(ksi,2)+pow(summ_ksi,2));
//     }
//     Array<double> Wm;Wm.SetSize(height->GetSize()*Ce.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//         {
//             Wm.Set(i*height->GetSize()+j,SNiP2_01_07_Formula6_Wn(Wo,Ce.Get(i),SNiP2_01_07_table6_K(height->Get(j),zone)));
//         }
//     }
//
//     double ro = YY;
//     double hi = H_max;
//     double Eps = SNiP2_01_07_Table9_Epsilon(ro, hi);
//     Array<double> Wp;Wp.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wp.Set(i*height->GetSize()+j,SNiP2_01_07_Formula9_Wp( Wm.Get(i*height->GetSize()+j), SNiP2_01_07_table7_Dzeta(H_max,zone),summ_ksi, Eps));
//     }
//
//     Array<double> Wsum;Wsum.SetSize(Wm.GetSize());
//     for(type_LLU i=0;i<Ce.GetSize();i++)
//     {
//         for(type_LLU j = 0;j<height->GetSize();j++)
//             Wsum.Set(i*height->GetSize()+j,Wm.Get(i*height->GetSize()+j)+Wp.Get(i*height->GetSize()+j));
//     }
//
//     // Find elements
//     // Add force
// }
//
//
//
//
//
//
//
// /*
// struct VeterPlastine
// {
//     double l ;
//     double k ;
//     double Wo;
//     double Ce;
//     double *Q;
//     double *M;
//     VeterPlastine(){;};
//     VeterPlastine(double _l , double  _k, double _Wo, double _Ce)
//     {
//          l = _l;
//          k = _k;
//          Wo = _Wo;
//          Ce = _Ce;
//     };
//     void Calculate(double *Q,double *M, bool out);
//     void Calculate(double *Q,double *M){Calculate( Q, M, false);};
// };
// ////////////////////////////////////////////
// /////    Расчет нагрузок на пластину   /////
// ////////////////////////////////////////////
// void VeterPlastine::Calculate(double *Q,double *M, bool out)
// {
// 	double Qvert,Mvert;
// 	double _Wn = Veter_BaseFunction::Wn( Wo, Ce, k, out);
// 	Qvert = _Wn*l/2;
// 	Mvert = _Wn*l*l/8.;
// 	if(out != false) // Вывод результатов //
// 	{
// 	    printf("\nРасчет ветровой нагрузки, действующую на пластину.\n");
// 		printf("Исходные данные:\n");
// 		printf("L = %.3f м - горизонтальный размер пластины.\n",l);
// 		printf("Решение:\n");
// 		printf("Q = Wn  x L/2 - погонная нагрузка на вертикальный стержень.\n");
// 		printf("Q = %.3f  x %.3f/2.\n",_Wn,l);
// 		printf("Q = %.2f Н/м.\n",Qvert);
// 		printf("M = Wn x L x L/8 - погонный момент на вертикальный стержень.\n");
// 		printf("M = %.3f x %.3f x %.3f/8.\n",_Wn,l,l);
// 		printf("M = %.2f Н*м/м.\n",Mvert);
// 	}
// 	*Q = Qvert;
// 	*M = Mvert;
// }
// ////////////////////////////////////////////
// /////    Расчет нагрузок на цилиндр    /////
// ////////////////////////////////////////////
// struct VeterCylinder
// {
//     double       *H; // Вектор отметок платин
//     type_LLU vn; // Количество вертикальных уровней
//     double        d; // Диаметр цилиндра
//     type_LLU gn; // Количество секций цилиндра
//     bool     middle; // Расположение первичной точки (true - на оси)
//     double       Wo; // Нормативное значение ветрового давления
//     zone          Z; // Ветровая зона
//     void printRawData(); // Вывод исходных данных
//     void Calculate();// Расчет задачи
// };
//
// void VeterCylinder::printRawData()
// {
// 	printf("\n\nИсходные данные:\n");
// 	printf("Номер отметки\tОтметка,м\n");
// 	type_LLU i;
// 	for (i=0 ; i<vn ; i++)
// 	    printf("%u\t%.3f\n",i,H[i]);
//     for (i=1 ; i<vn ; i++)
//         if(H[i] < H[i-1]) FATAL();
// 	printf("d  = %.3f м - Диаметр цилиндра, м.\n",d );
// 	printf("gn  = %u - Количество секций цилиндра, шт.\n",gn);
// 	printf("Wo = %.2f Па.\n",Wo);
// 	Veter_BaseFunction::printf_zone(Z);
// 	if(vn < 2 || gn < 6)FATAL();
// };
//
// void VeterCylinder::Calculate()
// {
//     type_LLU i;
//     for( i=0; i < (vn-1) ; i++ )
//     {
//         printf("\nРасчет уровня номер %u .\n", i+1 );
//         printf("\nРасчет производится по схеме 12б приложения 4 [__].\n");
//         double hi = H[i+1];
//         printf("Значение h1/d = %.3f.\n",H[vn-1]/d);
//         printf("Определение числа Рейнольдса(Re)\n");
//         double kz  = Veter_BaseFunction::Get_K(hi,Z);
//         printf("Значение коэффициента K для определения Re на высоте %.3f м составляет %.3f.\n",hi,kz);
//         printf("Расчет производится по схеме 12а приложения 4 [__].\n");
//         double _Re = Veter_BaseFunction::Get_Re(d,Wo,kz,true);
//         if ( _Re > 4e5) printf("Re = %.3f  > 400000 - Условие выполнено.\n",_Re);
//         else { FATAL(); printf("Re <= 400000.\n\n\n\n"); return; }
//         double K  = Veter_BaseFunction::Get_K(hi,Z);//k(H1,H2,Z);///ДОПУСКАЕТЬСЯ УМЕНЬШИТЬ ИСПОЛЬЗУЯ АППРОКСИМАЦИЮ
//         type_LLU j;
//         for(j=0;j<((type_LLU)(gn/2+0.6)); j++)
//         {
//             printf("\nУчасток №%u\n",j);
//             double angle1 = j*(360./gn);
//             double angle2 = (j+1)*(360./gn);
//             if(middle == false)
//             {
//                 angle1 -= (360./gn)/2.;
//                 angle2 -= (360./gn)/2.;
//             }
//             printf("Начальный угол - %.2f град.\n",angle1);
//             printf("Конечный  угол - %.2f град.\n",angle2);
//             double Ce1 = Veter_BaseFunction::Get_Ce1_Cxema12b(angle2,angle1,H[vn-1],d);
//             printf("Определение Wn:\n");
//             Veter_BaseFunction::Wn(Wo,Ce1,K,true);
//         }
//         if(middle == false)
//         {
//             printf("\nУчасток №%u\n",j);
//             double angle1 = j*(360./gn);
//             double angle2 = (j+1)*(360./gn);
//             if(middle == false)
//             {
//                 angle1 -= (360./gn)/2.;
//                 angle2 -= (360./gn)/2.;
//             }
//             printf("Начальный угол - %.2f град.\n",angle1);
//             printf("Конечный  угол - %.2f град.\n",angle2);
//             double Ce1 = Veter_BaseFunction::Get_Ce1_Cxema12b(angle2,angle1,H[vn-1],d);
//             printf("Определение Wn:\n");
//             Veter_BaseFunction::Wn(Wo,Ce1,K,true);
//         }
//     }
// }
// ////////////////////////////////////////////
// /////    Расчет нагрузок на пластину   /////
// ////////////////////////////////////////////
// struct VererStenka
// {
//     double H1;      // Верхняя граница пластин
//     double H2;      // Нижняя граница пластин
//     double *l;     // Горизонтальные размеры пластин
//     type_LLU n; // Количество пластин
//     double k;       // Коэффициент
//     double Wo;      // Нормативное значение ветрового давления
//     double Ce;      // Аэродинамический коэффициент
//     bool out;       // Вывод данных
//     void Calculate();// Расчет
//     void printRawData(); // Вывод исходных данных
// };
//
// void VererStenka::printRawData()
// {
// 	printf("\n\nИсходные данные:\n");
// 	printf("Номер стержня\tПоложение стержня\n");
// 	type_LLU i;
// 	for (i=0 ; i<n ; i++)
// 	    printf("%u\t%.3f\n",i+1,l[i]);
//     for (i=1 ; i<n ; i++)
//         if(l[i] < l[i-1]) FATAL();
// 	printf("k  = %.3f - коэффициент, м.\n",k );
// 	printf("n  = %u шт - Количество секций, шт.\n",n);
// 	printf("Wo = %.2f Па.\n",Wo);
// }
//
// void VererStenka::Calculate()
// {
//     type_LLU i;
// //    std::vector<VeterPlastine> Plastine(n-1);
//     Array<VeterPlastine> Plastine;
//     Plastine.SetSize(n-1);
//     for( i=0 ; i<(n-1); i++)
//     {
//          Plastine.Get(i).l = l[i+1]-l[i];
//          Plastine.Get(i).k = k;
//          Plastine.Get(i).Wo = Wo;
//          Plastine.Get(i).Ce = Ce;
//          double temp;
//          Plastine.Get(i).Calculate(&temp,&temp,out);
//     }
//     line();
//     printf("\nВывод данных полученных из расчетов:\n");
//     for(i=0;i<n;i++)
//     {
//         double Q=0;
//         double M=0;
//         if (i == 0)
//               Plastine.Get( 0 ).Calculate(&Q,&M,false);
//         else if (i == n)
//               Plastine.Get(n-1).Calculate(&Q,&M,false);
//         else
//         {
//             double q1,m1;
//             Plastine.Get(i-1).Calculate(&q1,&m1,false);
//             double q2,m2;
//             Plastine.Get( i ).Calculate(&q2,&m2,false);
//             Q = q1 + q2;
//             M = m1 - m2;
//         }
//         printf("для стержня: %i\n",i+1);
//         printf("Q = %.2f Н/м.\n",    Q);
//         printf("M = %.2f Н х м/м.\n",M);
//     }
// };
// //////////////////////////////////////////////
// ///// Расчет здания(камера               /////
// /////        радиации, конвекции)        /////
// //////////////////////////////////////////////
// struct VeterZdanie
// {
//     double        *X;
//     type_LLU nX;
//     double        *Y;
//     type_LLU nY;
//     double        *Z;
//     type_LLU nZ;
//     double        Wo;
//     zone        Zone;
//     bool         out;
//     void Calculate();
//     void Calculate_alternative();
// 	void printRawData();
// };
// void VeterZdanie::printRawData()
// {
//     type_LLU i=0;
//     printf("Расчет ветровой нагрузки на здание по схеме 2 приложения 4 СНиП 2.01.07-85*.\n");
//     printf("Исходные данные.\n");
//     Veter_BaseFunction::printf_zone( Zone );
//     if(nX < 2 || nY < 2 || nZ < 2) FATAL();
//     printf("Значение Wo = %.1f Па.\n",Wo);
//     printf("Количество вертикальных участков - %u.\n",(nZ-1));
//     printf("Количество горизонтальных участков по оси X - %u.\n",(nX-1));
//     printf("Количество горизонтальных участков по оси Y - %u.\n",(nY-1));
//     printf("Номер отметки\tОтметка, м\n");
//     for( i=0; i<nZ ; i++ )
//     {
//         printf("%u\t%.3f\n",i,Z[i]);
//     }
//     for( i=1; i<nZ ; i++ )
//          if(Z[i]<Z[i-1])FATAL();
//     printf("Номер точки по X\tПоложение, м\n");
//     for( i=0; i<nX ; i++ )
//     {
//         printf("%u\t%.3f\n",i,X[i]);
//     }
//     for( i=1; i<nX ; i++ )
//          if(X[i]<X[i-1])FATAL();
//     printf("Номер точки по Y\tПоложение, м\n");
//     for( i=0; i<nY ; i++ )
//     {
//         printf("%u\t%.3f\n",i,Y[i]);
//     }
//     for( i=1; i<nY ; i++ )
//          if(Y[i]<Y[i-1])FATAL();
// }
//
// void VeterZdanie::Calculate()
// {
//     type_LLU i;
//     for(type_LLU j=0; j<2 ; j++)
//     {
//         if(j==0)printf("\n\nОпределение ветровой нагрузки при действии ветра по оси X\n");
//         else
//         if(j==1)printf("\n\nОпределение ветровой нагрузки при действии ветра по оси Y\n");
//         else FATAL();
//         for( i=0 ; i<(nZ-1) ; i++)
//         {
//             double *Front;  // Передняя стенка
//             type_LLU nFront;
//             double *Sidebar;// Боковая  стенка
//             type_LLU nSidebar;
//             if(j == 0 )
//             {
//                 Front    = Y;
//                 nFront   = nY;
//                 Sidebar  = X;
//                 nSidebar = nX;
//             }
//             else
//             {
//                 Front    = X;
//                 nFront   = nX;
//                 Sidebar  = Y;
//                 nSidebar = nY;
//             }
//             printf("Уровень номер %u от отметки %.3f до отметки %.3f\n",i,Z[i],Z[i+1]);
//             double h1 = Z[i+1] - Z[i];
//             double k  = Veter_BaseFunction::Get_K(Z[i+1],Zone);
//             printf("Значение коэффициента K высоте %.3f м составляет %.3f.\n",Z[i+1],k);
//             /////////
//             printf("\nОпределение значений ветровой нагрузки с наветренной стороны\n");
//             VererStenka Stenka1;
//             Stenka1.Ce = 0.8;
//             Stenka1.H1 = Z[i];
//             Stenka1.H2 = Z[i+1];
//             Stenka1.k  = k;
//             Stenka1.n  = nFront;
//             Stenka1.out= out;
//             Stenka1.Wo = Wo;
//             Stenka1.l  = Front;
//             if(out == true)Stenka1.printRawData();
//             Stenka1.Calculate();
//             /////////
//             printf("\nОпределение значений ветровой нагрузки с подветренней стороны\n");
//             double Ce3 = Veter_BaseFunction::Get_Ce3_Cxema2(h1,Sidebar[nSidebar-1]-Sidebar[0],Front[nFront-1]-Front[0]);
//             Stenka1.Ce = Ce3;
//             if(out == true)Stenka1.printRawData();
//             Stenka1.Calculate();
//             /////////
//             printf("\nОпределение значений ветровой нагрузки с торцевой стороны\n");
//             VererStenka Stenka2;
//             Stenka2.Ce = Ce3;
//             Stenka2.H1 = Z[i];
//             Stenka2.H2 = Z[i+1];
//             Stenka2.k  = k;
//             Stenka2.n  = nSidebar;
//             Stenka2.out= out;
//             Stenka2.Wo = Wo;
//             Stenka2.l  = Sidebar;
//             if(out == true)Stenka2.printRawData();
//             Stenka2.Calculate();
//         }
//     }
// }
//
// void VeterZdanie::Calculate_alternative()
// {
//     type_LLU i;
//     for(type_LLU j=0; j<2 ; j++)
//     {
// 		line();
//         if(j==0)printf("\n\nОпределение ветровой нагрузки при действии ветра по оси X\n");
//         else
//         if(j==1)printf("\n\nОпределение ветровой нагрузки при действии ветра по оси Y\n");
//         else FATAL();
//         for( i=0 ; i<(nZ-1) ; i++)
//         {
//             double *Front;  // Передняя стенка
//             type_LLU nFront;
//             double *Sidebar;// Боковая  стенка
//             type_LLU nSidebar;
//             if(j == 0 )
//             {
//                 Front    = Y;
//                 nFront   = nY;
//                 Sidebar  = X;
//                 nSidebar = nX;
//             }
//             else
//             {
//                 Front    = X;
//                 nFront   = nX;
//                 Sidebar  = Y;
//                 nSidebar = nY;
//             }
//             printf("Уровень номер %u от отметки %.3f до отметки %.3f\n",i,Z[i],Z[i+1]);
//             double h1 = Z[i+1] - Z[i];
//             double k  = Veter_BaseFunction::Get_K(Z[i+1],Zone);
//             printf("Значение коэффициента K на высоте %.3f м составляет %.3f.\n",Z[i+1],k);
//             /////////
//             printf("\nОпределение значений ветровой нагрузки с наветренной стороны\n");
//             Veter_BaseFunction::Wn(Wo,0.8,k,true);
//             /////////
//             printf("\nОпределение значений ветровой нагрузки с подветренней стороны\n");
//             double Ce3 = Veter_BaseFunction::Get_Ce3_Cxema2(h1,Sidebar[nSidebar-1]-Sidebar[0],Front[nFront-1]-Front[0]);
//             Veter_BaseFunction::Wn(Wo,Ce3,k,true);
//             /////////
//             printf("\nОпределение значений ветровой нагрузки с торцевой стороны\n");
//             Veter_BaseFunction::Wn(Wo,Ce3,k,true);
//         }
//     }
// }
// //////////////////////////////////////////////
// //// Расчет горизонтальных труб           ////
// //////////////////////////////////////////////
// struct VeterTrube_Gorizontal
// {
//     double D;
//     double H;
//     double l;
//     double Wo;
//     zone Zone;
//     void Calculate();
// 	void printRawData();
// };
//
// void VeterTrube_Gorizontal::printRawData()
// {
//     printf("\n\nРасчет горизонтальной трубы по схеме 13 приложения 4 СНиП 2.01.07-85*\n");
//     printf("Исходные данные\n");
//     printf("D = %.3f м - диаметр трубы\n",D);
//     printf("l = %.3f м - горизонтальная длина трубы\n",l);
//     printf("Н = %.3f м - отметка высоты оси трубы\n",H);
//     Veter_BaseFunction::printf_zone( Zone );
//     if(D < 0 || l < 0 || H < 0) FATAL();
//     printf("Значение Wo = %.1f Па\n",Wo);
// }
//
// void VeterTrube_Gorizontal::Calculate()
// {
//     printf("Определение значения lambda\n");
//     double lambda = D/l;
//     printf("lambda = D/l = %.3f / %.3f\n",D,l);
//     printf("lambda = %.3f\n", lambda);
//     printf("Определение значения lambda_e\n");
//     printf("lambda_e = lambda = %.3f\n", lambda);
//     printf("Определение значения коэффициента k по таблице 1 схемы 13\n");
//     double k = Veter_BaseFunction::Get_K_Cxema13(lambda);
//     printf("k = %.3f\n", k);
//     printf("Определение значения K (K_6) по таблице 6 СНиП 2.01.07-85*\n");
//     double K_6 = Veter_BaseFunction::Get_K(H,Zone);
//     printf("K_6 = %.3f\n", K_6);
//     printf("Определение числа Рейнольдса по схемы 12а\n");
//     double Re = Veter_BaseFunction::Get_Re(D,Wo,K_6,true);
//     printf("Определение значения аэродинамического коэффициента Ce_x_bez по графику схемы 14\n");
//     double Ce_x_bez = Veter_BaseFunction::Get_Ce_x_bez(0.001/D, Re);
//     printf("Ce_x_bez = %.3f\n", Ce_x_bez);
//     printf("Определение значения аэродинамического коэффициента Ce\n");
//     double Ce = Ce_x_bez * K_6;
//     printf("Ce = Ce_x_bez * K_6 = %.3f * %.3f\n", Ce_x_bez, K_6);
//     printf("Ce = %.3f\n", Ce);
//
//     double wn = Veter_BaseFunction::Wn(Wo,Ce,K_6,true);
//     printf("Определение нагрузки на стержень\n");
//     double Qn = Veter_BaseFunction::Qn(wn,D);
//     printf("Qn = %.3f Н/м\n", Qn);
//
// }
// */
//
//
//
//
// //class PLATE
// //{
// //    double R, THK;
// //    double a, b;
// //    PLATE(){R = THK = a = b = 0};
// //    double SigmaCritical(bool OUT)
// //    {
// //
// //    }
// //}
//
// //double ShellWithStiffeners(
// //     double Sigma,
// //     double Sigma_Max,
// //     double Shell_D,
// //     double Shell_THK,
// //     double Shell_H,
// //
// //     double Stiff_Vert_Area,
// //     double Stiff_Vert_Jx,
// //     double Stiff_Vert_Ymax,
// //     type_LLU Stiff_Vert_Items,
// //
// //     type_LLU Stiff_Horiz_Items
// //     )
// //{
// //    double factor = 0;
// //    //double E = 2.06e11;
// //
// //    printf("Check 1\n");
// //    printf("Sigma = %.1fMPa < Sigma_Max = %.1fMPa\n",Sigma*1e-6,Sigma_Max*1e-6);
// //    Printf_CALC(Sigma/Sigma_Max,factor,true);
// //
// //    printf("Check 2\n");
// //    double Lplate = M_PI*Shell_D/Stiff_Vert_Items;
// //    if(SNiP_II_23_p8_7(Lplate, Shell_D/2, Shell_THK, Sigma, Sigma_Max,true))
// //    {
// //        //SNiP_II_23_p8_7(Lplate, Shell_D/2, Shell_THK, Sigma, Sigma_Max,true);
// //    }
// //    else
// //    {
// //        double SigmaCr_SNiP;
// //        SNiP_II_23_p8_5(Shell_D/2,Shell_THK,Sigma,Sigma_Max,SigmaCr_SNiP);
// //        Printf_CALC(Sigma/SigmaCr_SNiP,factor,true);
// //    }
// //
// //
// //    return factor;
// //}
