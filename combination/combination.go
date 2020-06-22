package combination

// ////////////////////////////////////
// // COMBINATIONS
// enum LOAD_SNIP{LOAD_SNIP_CONST, LOAD_SNIP_SHORT_TIME, LOAD_SNIP_LONG_TIME, LOAD_SNIP_EXTREMAL,LOAD_SNIP_SNOW};
// enum LOADS{ LOAD_STEEL=0, LOAD_REFRECTORY,
//             LOAD_COIL , LOAD_FLUID     ,
//             LOAD_SEISMIC, LOAD_WIND,
//             LOAD_SNOW  ,LOAD_LIVELOAD,
//             LOAD_ON_PIPING,
//             LOAD_ON_REACTIONS,
//             LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST};
// void Printf(LOAD_SNIP ls)
// {
//     switch(ls)
//     {
//         case LOAD_SNIP_CONST        : printf("Constant   load SNiP2.01.07\n"); break;
//         case LOAD_SNIP_SHORT_TIME   : printf("Short time load SNiP2.01.07\n"); break;
//         case LOAD_SNIP_LONG_TIME    : printf("Long time  load SNiP2.01.07\n"); break;
//         case LOAD_SNIP_EXTREMAL     : printf("Extremal   load SNiP2.01.07\n"); break;
//         case LOAD_SNIP_SNOW         : printf("Snow       load SNiP2.01.07\n"); break;
//         default:
//             print_name("ERROR in printf LOAD_SNIP");
//             printf("ls=%u\n",ls);
//             FATAL();
//     }
// }
//
// void Printf(LOADS l)
// {
//     switch(l)
//     {
//         case LOAD_STEEL        : printf("LOAD_STEEL"); break;
//         case LOAD_REFRECTORY   : printf("LOAD_REFRECTORY"); break;
//         case LOAD_COIL         : printf("LOAD_COIL"); break;
//         case LOAD_FLUID        : printf("LOAD_FLUID"); break;
//         case LOAD_SEISMIC      : printf("LOAD_SEISMIC"); break;
//         case LOAD_WIND         : printf("LOAD_WIND"); break;
//         case LOAD_LIVELOAD     : printf("LOAD_LIVELOAD"); break;
//         case LOAD_SNOW         : printf("LOAD_SNOW"); break;
//         case LOAD_ON_PIPING    : printf("LOAD_ON_PIPING"); break;
//         case LOAD_ON_REACTIONS : printf("LOAD_ON_REACTIONS"); break;
//         default:
//             print_name("ERROR in ConvertLOADStoLOAD_SNIP");
//             printf("l=%u\n",l);
//             FATAL();
//     }
// }
//
// bool ConvertLOADStoLOAD_SNIP(LOADS l, LOAD_SNIP &ls)
// {
//     switch(l)
//     {
//         case LOAD_STEEL                 : ls = LOAD_SNIP_CONST;      break;
//         case LOAD_REFRECTORY            : ls = LOAD_SNIP_CONST;      break;
//         case LOAD_COIL                  : ls = LOAD_SNIP_CONST;      break;
//         case LOAD_FLUID                 : ls = LOAD_SNIP_CONST;      break;
//         case LOAD_SEISMIC               : ls = LOAD_SNIP_EXTREMAL  ; break;
//         case LOAD_WIND                  : ls = LOAD_SNIP_SHORT_TIME; break;
//         case LOAD_LIVELOAD              : ls = LOAD_SNIP_SHORT_TIME; break;
//         case LOAD_SNOW                  : ls = LOAD_SNIP_SNOW;       break;
//         case LOAD_ON_PIPING             : ls = LOAD_SNIP_CONST;      break;
//         case LOAD_ON_REACTIONS          : ls = LOAD_SNIP_CONST;      break;
//         default:
//             print_name("ERROR in ConvertLOADStoLOAD_SNIP");
//             printf("l=%u ls=%u\n",l,ls);
//             FATAL();
//     }
//     return true;
// }
//
// struct LOAD
// {
//     float factor;
//     LOADS type_LOAD;
//     type_LLU NumberCase;
//     char *name;
//     void Printf();
//     void Printf();
//     void prn();
//     LOAD & operator= (const LOAD & param)
//     {
//         factor          = param.factor;
//         type_LOAD       = param.type_LOAD;
//         NumberCase = param.NumberCase;
//         name = new char[20];
//         strcpy(name,param.name);
//         return *this;
//     }
//     LOAD()
//     {
//         factor          = 0;
//     	NumberCase = 0;
//     	type_LOAD  = LOAD_STEEL;
//     	name = new char[20];
//     };
//     LOAD(double _factor, type_LLU _NumberCase,LOADS _type_LOAD,const char *_name)
//     {Init(_factor, _NumberCase, _type_LOAD, _name);};
//     void Init(double _factor, type_LLU _NumberCase,LOADS _type_LOAD,const char *_name)
//     {
//
//         factor          = _factor;
//         NumberCase = _NumberCase;
//         type_LOAD       = _type_LOAD;
//         name = new char[20];
//         strcpy(name,_name);
//     }
//     bool operator== (const LOAD & param)
//     {
//         bool DEBUG = false;//true;//
//         if(type_LOAD       != param.type_LOAD      )
//             {if(DEBUG)printf("Diff LOAD+1");return false;}
//         if(fabs(factor - param.factor)>1e-6)
//             {if(DEBUG)printf("Diff LOAD+2");return false;}
//         return true;
//     }
// };
// void LOAD::Printf()
// {
//     LOAD_SNIP ls;
//     ConvertLOADStoLOAD_SNIP(type_LOAD, ls);
//     printf("%.3f\t\"%20s\"\t",factor,name);::Printf(type_LOAD);printf("\t");::Printf(ls);//printf("\n");
// }
// void LOAD::Printf()
// {
//     if(NumberCase<= 0 || fabs(factor)< 1e-6)
//     { print_name("Error in LOAD::Printf()"); printf("%5u %e\n",NumberCase,factor);FATAL(); }
//     printf(" %u %.3f ",NumberCase,factor);
// }
// void LOAD::prn()
// {
//     if(NumberCase<= 0 || fabs(factor)< 1e-6)
//     { print_name("Error in LOAD::prn()"); printf("%5u %e\n",NumberCase,factor);FATAL(); }
//     printf("%.3fx\"%s\" ",factor,name);
// }
//
// enum COMBINATION_TYPE{COMBINATION_TYPE_SHORT,COMBINATION_TYPE_FULL};
// /*
// struct COMBINATION
// {
//     Array<LOAD>        loads;
//     Array<COMBINATION>  comb;
//     type_LLU NumberCase;
//     void Printf();
//     void Printf();
//     void Calculate(Array<LOAD> &Combinations,COMBINATION_TYPE CT);
//     ~COMBINATION(){loads.Delete();comb.Delete();}
// };
//
// void COMBINATION::Printf()
// {
//     type_LLU number_combination = 0;
//     for(type_LLU j=0;j<comb.GetSize();j++)
//     {
//         number_combination++;
//         printf("%5u) ",number_combination);
//         for(type_LLU i=0;i<comb.Get(j).loads.GetSize();i++)
//         {
//             if(i!=0)printf("+ ");
//             comb.Get(j).loads.Get(i).prn();
//         }
//         printf("\n");
//     }
//     if(comb.GetSize() == 0)
//     {
//         for(type_LLU i=0;i<loads.GetSize();i++)
//         {
//             if(i!=0)printf("+ ");
//             loads.Get(i).prn();
//         }
//         printf("\n");
//     }
//     //printf("%u %u\n",comb.GetSize(),loads.GetSize());
// }
//
// void COMBINATION::Printf()
// {
//     // LOAD COMB 100 TEST
//     // 2 1.0 3 1.0
//     type_LLU number_combination = 0;
//     for(type_LLU j=0;j<comb.GetSize();j++)
//     {
//         number_combination++;
//         printf("LOAD COMB %5u Combination%u\n",NumberCase+number_combination,NumberCase+number_combination);
//         for(type_LLU i=0;i<comb.Get(j).loads.GetSize();i++)
//         {
//             comb.Get(j).loads.Get(i).Printf();
//         }
//         printf("\n");
//     }
// }
// */
// struct COMBINATION
// {
//     Array<LOAD>        loads;
//     type_LLU NumberCase;
//     void Printf();
//     void Printf();
//     ~COMBINATION(){loads.Delete();}
//     bool operator== (COMBINATION & param)
//     {
//         bool DEBUG = false;//true;//
//         if(loads.GetSize() != param.loads.GetSize())
//         {
//             if(DEBUG)printf("Diff 1");
//             return false;
//         }
//         for(type_LLU i = 0;i<loads.GetSize();i++)
//             if(!(loads[i] == param.loads[i]))
//                 {
//                     if(DEBUG)printf("Diff 2{%u}",i);
//                     return false;
//                 }
//         if(DEBUG)printf("some");
//         return true;
//     }
//     COMBINATION & operator= (const COMBINATION & param)
//     {
//         loads           = param.loads;
//         NumberCase = param.NumberCase;
//         return *this;
//     }
// };
//
// void COMBINATION::Printf()
// {
//     for(type_LLU i=0;i<loads.GetSize();i++)
//     {
//         if(i!=0)printf("+ ");
//         loads.Get(i).prn();
//     }
//     printf("\n");
// }
//
// void COMBINATION::Printf()
// {
//     //*LOAD COMB 100 TEST
//     //*2 1.0 3 1.0
//     printf("LOAD COMB %5u Combination%u\n",NumberCase,NumberCase);
//     for(type_LLU i=0;i<loads.GetSize();i++)
//     {
//         loads.Get(i).Printf();
//     }
//     printf("\n");
// }
//
// void COMBINATION_Calculate(Array<LOAD> &Combinations,COMBINATION_TYPE Comb_Type,
//                            type_LLU NumberCase, Array<COMBINATION>  &comb)
// {
//     bool DEBUG = false;//true;//
//     if(DEBUG)
//         for(type_LLU i = 0;i<Combinations.GetSize();i++)
//             Combinations.Get(i).Printf();
//     type_LLU MaxComb = (type_LLU)pow(2,Combinations.GetSize());
//     if(DEBUG)printf("Max combinations is %u\n",MaxComb);
//     Array<COMBINATION> CombWithoutSNiP;
//     COMBINATION *tmpC;
//     type_LLU MAX_SIZE_COMB = MaxComb;
//     /// Calc full number of load ///
//     type_LLU *NumberALL = new type_LLU[LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST];
//     for(type_LLU j=0;j<LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST;j++)NumberALL[j]=0;
//     for(type_LLU j=0;j<Combinations.GetSize();j++)
//             NumberALL[Combinations.Get(j).type_LOAD]++;
//     if(DEBUG)
//         for(type_LLU j=0;j<LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST;j++)
//             {Printf((LOADS)j);printf("\tNumberALL[%2u] = %u\n",j,NumberALL[j]);}
//     /// End calc ///
//     for(type_LLU GNN = 0;GNN<2;GNN++)
//     {
//         // GNN=1 - calculate comb
//         // GNN=2 - create comb
//         type_LLU SIZE_COMB = 0;
//         for(type_LLU i=0;i<MaxComb;i++)
//         {
//             /// Convert from decimal to bit format ///
//             type_LLU tmp = i;
//             Array<short> num;
//             for(type_LLU j=0;j<Combinations.GetSize();j++)
//             {
//                 if(tmp == 1) {num.Add(1);break;}
//                 else if(tmp == 0){num.Add(0);break;}
//                 else
//                 if((double)tmp/2 == (type_LLU)tmp/2){num.Add(0);tmp/=2;}
//                 else {tmp=(tmp-1)/2; num.Add(1);}
//             }
//             tmp = num.GetSize();
//             for(type_LLU j=0;j<(Combinations.GetSize()-tmp);j++)
//                 num.Add(0);
//             /*if(DEBUG)
//             {
//                 printf("%6u\t",i);
//                 for(type_LLU i=num.GetSize();i>0;i--)
//                     printf("%d",num[i-1]);
//                 printf("\n");
//             }*/
//             /// End converting ///
//
//             /// Calculate number of loads ///
//             type_LLU *Number = new type_LLU[LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST];
//             for(type_LLU j=0;j<LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST;j++)Number[j]=0;
//             for(type_LLU j=0;j<num.GetSize();j++)
//                 if(num.Get(j) == 1)
//                     Number[Combinations.Get(j).type_LOAD]++;
//             /// End calculation number of loads ///
//
//             //if(GNN != 0 && MAX_SIZE_COMB == SIZE_COMB) break;
//
//             /// FIRST FILTER ///
//             if(!(
//                Number[LOAD_REFRECTORY]   == 1 && Number[LOAD_STEEL]        == 1 &&
//                (NumberALL[LOAD_COIL]         == 0 || Number[LOAD_COIL]         == 1) &&
//                (NumberALL[LOAD_ON_REACTIONS] == 0 || Number[LOAD_ON_REACTIONS] == 1) &&
//                Number[LOAD_SNOW]         <  2 && Number[LOAD_LIVELOAD]     <  2 &&
//                Number[LOAD_FLUID]        <  2 && Number[LOAD_ON_PIPING   ] <  2 &&
//                ((Number[LOAD_WIND]        ==0 && Number[LOAD_SEISMIC]     ==0) ||
//                 (Number[LOAD_WIND]        ==1 && Number[LOAD_SEISMIC]     ==0) ||
//                 (Number[LOAD_WIND]        ==0 && Number[LOAD_SEISMIC]     ==1))
//               ))
//             {
//                 delete []Number;
//                 num.Delete();
//                 continue;
//             }
//
//             /// TYPICAL COMBINATION///
//             if( (Comb_Type == COMBINATION_TYPE_FULL) ||
//                 (Comb_Type != COMBINATION_TYPE_FULL      &&
//                      (Number[LOAD_WIND]      == 1 || Number[LOAD_SEISMIC]     ==1) &&
//                       Number[LOAD_ON_PIPING] == 1 &&
//                       Number[LOAD_SNOW]      == 1 &&
//                       Number[LOAD_LIVELOAD]  == 1)
//               )
//             {
//                 if(GNN != 0)
//                 {
//                     tmpC = new COMBINATION;
//                     for(type_LLU j=0;j<num.GetSize();j++)
//                         if(num.Get(j) == 1)
//                             tmpC->loads.Add(Combinations.Get(j));
//                     CombWithoutSNiP.Set(SIZE_COMB,*tmpC);
//                     delete tmpC;
//                 }
//                 SIZE_COMB++;
//             }
//             /// JN LEROUX Cases ///
//
//             if(
//                    //Number[LOAD_REFRECTORY]  == 1 &&
//                    //Number[LOAD_STEEL]       == 1 &&
//                    //Number[LOAD_COIL]        == 1 &&
//                    ((Number[LOAD_WIND]        == 1 && Number[LOAD_SEISMIC]     == 0)  ||
//                     (Number[LOAD_WIND]        == 0 && Number[LOAD_SEISMIC]     == 1)) &&
//                    Number[LOAD_SNOW]        == 0 &&
//                    Number[LOAD_LIVELOAD]    == 0 &&
//                    Number[LOAD_FLUID]       == 0 &&
//                    //Number[LOAD_ON_PIPING   ]== 1 &&
//                    Number[LOAD_ON_REACTIONS]== 1)
//             {
//                 if(GNN != 0)
//                 {
//                     double ff = 0.0;
//                     tmpC = new COMBINATION;
//                     double LerouxFactor;
//                     if(Number[LOAD_WIND] == 1) LerouxFactor = 0.9;
//                     else LerouxFactor = 1.0;
//                     for(type_LLU j=0;j<num.GetSize();j++)
//                         if(num.Get(j) == 1)
//                         {
//                             if(Combinations.Get(j).type_LOAD == LOAD_REFRECTORY ||
//                                Combinations.Get(j).type_LOAD == LOAD_STEEL      ||
//                                Combinations.Get(j).type_LOAD == LOAD_COIL       )
//                             {
//                                 ff = Combinations.Get(j).factor;
//                                 if(ff != LerouxFactor) Combinations.Get(j).factor = LerouxFactor;
//                             }
//                             tmpC->loads.Add(Combinations.Get(j));
//                             if(Combinations.Get(j).type_LOAD == LOAD_REFRECTORY ||
//                                Combinations.Get(j).type_LOAD == LOAD_STEEL      ||
//                                Combinations.Get(j).type_LOAD == LOAD_COIL        )
//                                     Combinations.Get(j).factor = ff;
//                         }
//                     //tmp.Printf();
//                     CombWithoutSNiP.Set(SIZE_COMB,*tmpC);
//                     delete tmpC;
//                 }
//                 SIZE_COMB++;
//             }
//             delete []Number;
//             num.Delete();
//         }
//         if(GNN == 0)
//         {
//             CombWithoutSNiP.SetSize(SIZE_COMB);
//             MAX_SIZE_COMB = SIZE_COMB;
//             if(DEBUG)printf("Max combinations is %u\n",MaxComb);
//             if(DEBUG)printf("Max combinations is %u\n",MAX_SIZE_COMB);
//         }
//     }
//     /// Create loads with SNiP ///
//     /// Loads for deformation ///
//     Array<COMBINATION> CombWithoutSNiPforDeformation;
//     CombWithoutSNiPforDeformation = CombWithoutSNiP;
// //    CombWithoutSNiPforDeformation = CombWithoutSNiP;
// //    for(type_LLU j=0;j<CombWithoutSNiPforDeformation.GetSize();j++)
// //    {
// //        for(type_LLU i=0;i<CombWithoutSNiPforDeformation.Get(j).loads.GetSize();i++)
// //            CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 1;
// //    }
// //    comb = CombWithoutSNiPforDeformation;
//
//     /// Load from SNiPs ///
//     Array<COMBINATION> CombWithSNiP;
//     CombWithSNiP = CombWithoutSNiP;
//     for(type_LLU j=0;j<CombWithoutSNiP.GetSize();j++)
//     {
//         //if(DEBUG)printf("[%u/%u]\n",j,CombWithoutSNiP.GetSize());
//         Array<LOAD_SNIP> ls;
//         ls.SetSize(CombWithoutSNiP.Get(j).loads.GetSize());
//         type_LLU num_snow = 100000;
//         type_LLU num_extr = 100000;
//         for(type_LLU i=0;i<CombWithSNiP.Get(j).loads.GetSize();i++)
//         {
//             ConvertLOADStoLOAD_SNIP(CombWithSNiP.Get(j).loads.Get(i).type_LOAD,ls.Get(i));
//             if(CombWithoutSNiP.Get(j).loads.Get(i).type_LOAD != LOAD_SNOW)
//                 num_snow = i;
//             if(CombWithoutSNiP.Get(j).loads.Get(i).type_LOAD != LOAD_SEISMIC)
//                 num_extr = i;
//         }
//         /// Calculate number of loads ///
//         type_LLU *Number = new type_LLU[LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST];
//         for(type_LLU k=0;k<LOAD_ONLY_FOR_NUMBER_OF_LOADS_LAST;k++)Number[k]=0;
//         for(type_LLU k=0;k<ls.GetSize();k++)
//                 Number[ls.Get(k)]++;
//         if(Number[LOAD_SNIP_EXTREMAL] == 0)
//         {
//             /// SNiP 2.01.07 ///
//             if(Number[LOAD_SNIP_LONG_TIME]+Number[LOAD_SNIP_SHORT_TIME]+Number[LOAD_SNIP_SNOW] == 1)
//             {
//                 for(type_LLU i=0;i<CombWithSNiP.Get(j).loads.GetSize();i++)
//                 {
//                     CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 1.0;
//                     CombWithSNiP.Get(j).loads.Get(i).factor *= 1.00;
//                 }
//             }
//             else
//             {
//                 for(type_LLU i=0;i<CombWithSNiP.Get(j).loads.GetSize();i++)
//                 {
//                     if(ls.Get(i) == LOAD_SNIP_LONG_TIME)
//                     {
//                         CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 0.95;
//                         CombWithSNiP.Get(j).loads.Get(i).factor *= 0.95;
//                     }
//                     else if(ls.Get(i) == LOAD_SNIP_SHORT_TIME || ls.Get(i) == LOAD_SNIP_SNOW)
//                     {
//                         CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 0.90;
//                         CombWithSNiP.Get(j).loads.Get(i).factor *= 0.90;
//                     }
//                     else
//                         CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 1.00;
//                 }
//             }
//         }
//         else
//         {
//             /// SNiP II-7 ///
//             for(type_LLU i=0;i<CombWithSNiP.Get(j).loads.GetSize();i++)
//             {
//                 if(ls.Get(i) == LOAD_SNIP_LONG_TIME)
//                 {
//                     CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 0.8;
//                     CombWithSNiP.Get(j).loads.Get(i).factor *= 0.8;
//                 }
//                 else if(ls.Get(i) == LOAD_SNIP_SHORT_TIME || ls.Get(i) == LOAD_SNIP_SNOW)
//                 {
//                     CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 0.5;
//                     CombWithSNiP.Get(j).loads.Get(i).factor *= 0.5;
//                 }
//                 else if(ls.Get(i) == LOAD_SNIP_CONST)
//                 {
//                     CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 0.9;
//                     CombWithSNiP.Get(j).loads.Get(i).factor *= 0.9;
//                 }
//                 else if(ls.Get(i) == LOAD_SNIP_EXTREMAL)
//                 {
//                     CombWithoutSNiPforDeformation.Get(j).loads.Get(i).factor = 1.0;
//                     CombWithSNiP.Get(j).loads.Get(i).factor *= 1.0;
//                 }
//                 else{::Printf(CombWithSNiP.Get(j).loads.Get(i).type_LOAD);print_name("Error in /// SNiP II-7 ///\n");FATAL();}
//             }
//         }
//         ls.Delete();
//     }
//     type_LLU Nmm = 0;
//     if(false)//IF I WANT ADD DEFORMATION THEN (CT == COMBINATION_TYPE_FULL)
//     {
//         Nmm = CombWithoutSNiPforDeformation.GetSize();
//         comb.SetSize(CombWithoutSNiPforDeformation.GetSize()+CombWithSNiP.GetSize());
//         for(type_LLU i = 0;i<CombWithoutSNiPforDeformation.GetSize();i++)
//             comb.Set(i,CombWithoutSNiPforDeformation.Get(i));
//     }
//     else comb.SetSize(CombWithSNiP.GetSize());
//     for(type_LLU i = 0;i<CombWithSNiP.GetSize();i++)
//         comb.Set(i+Nmm,CombWithSNiP.Get(i));
//     /// Checking
//     comb.DeleteEqual();
//     if(comb.GetSize() <= 1)
//     { print_name("Error in COMBINATION_Calculate1"); FATAL(); }
//     for(type_LLU i=0;i<comb.GetSize();i++)
//     {
//         if(comb[i].loads.GetSize() <= 1)
//         { print_name("Error in COMBINATION_Calculate2"); FATAL(); }
//         for(type_LLU j=0;j<comb[i].loads.GetSize();j++)
//         {
//             LOAD L = comb[i].loads[j];
//             if(L.NumberCase<= 0 || fabs(L.factor)< 1e-6)
//             { print_name("Error in COMBINATION_Calculate3"); FATAL(); }
//         }
//     }
//     for(type_LLU i=0;i<comb.GetSize();i++)
//     {
//         comb[i].NumberCase = NumberCase+i;
//     }
//     if(DEBUG)printf("Size CombWithoutSNiP is %u\n",CombWithoutSNiP.GetSize());
//     if(DEBUG)printf("Size comb.GetSize()  is %u\n",comb.GetSize());
//     //for(type_LLU i=0;i<comb.GetSize();i++)
//     //    comb[i].Printf();
//     /// Deleting
// //    CombWithoutSNiP.Delete();
// //    CombWithoutSNiPforDeformation.Delete();
// //    CombWithSNiP.Delete();
// }
