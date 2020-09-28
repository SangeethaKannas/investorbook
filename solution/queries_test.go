package main

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestGetFirstDegree(t *testing.T) {
	a := App{}
	a.InitializeAndRun()
	defer a.DBClose()
	type args struct {
		db         *sql.DB
		investorId string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"test1", args{a.DB, "1"}, []int{9944,5847,819,3474,5067,4548,5869,5326,5596,4012,3907,7530,3057,8635,8921,2480,810,8313,5560,9177,3883,4684,8067,4306,3143,6926,7809,9257,6359,105,2508,3980,6564,4685,688,40,7623,5296,6908,786,3439,5466,817,2214,9287,632,6410,4068,9357,8263,4888,597,4887,4492,8095,8408,7416,6423,8246,9612,2316,2353,4038,6278,6056,9054,5718,1122,2405,9183,9504,4698,3784,8920,2138,2330,9215,7866,5378,3555,1068,6434,9530,6660,8041,6785,5360,8945,160,8091,5027,6540,3212,673,7470,2827,650,4141,4409,3057,5816,6596,9094,4473,2595,274,8508,216,5634,7962,9230,2685,9108,3268,3076,3186,963,8718,7644,1392,6939,8509,2838,5952,2237,3446,1402,8496,6009,7306,6045,4079,5027,9835,329,6664,1653,9170,3833,9882,2832,1392,4905,9854,8718,7066,6002,7908,3478,4576,2606,9630,8922,3157,7065,8437,2375,4798,5384,1807,1564,88,6925,1342,5805,3085,7621,271,9480,8661,7140,3139,5019,9470,808,1753,7647,6666,3582,4266,330,7797,5472,4635,640,3825,144,6326,1713,2238,8508,6531,4567,6856,5969,6627,2313,9256,602,6499,5324,2718,6836,4620,9367,6737,7308,9688,9640,3644,5354,5617,6082,760,9392,1339,608,5545,180,1080,2595,3767,7564,5130,8417,896,6273,2901,5331,7393,6997,6905,128,148,8560,9624,9686,689,4162,4488,4387,6670,6489,7487,2682,3842,1572,2938,4374,5961,3458,7520,4225,3298,2488,3938,6986,7318,3031,3387,6520,3504,1761,3276,7018,8518,9521,3593,6860,6400,3917,4007,1624,7705,7405,31219,352,22702,42748,18502,32410,29347,47907,10357,17158,36503,49627,28851,20881,45101,23143,28635,48383,37488,38304,48921,20001,32480,17809,36926,49257,33980,46359,20105,32508,10688,46564,14685,40001,30040,36908,47623,12534,45296,14887,44068,19357,23246,28263,18408,27416,44492,48095,39054,34888,48246,36056,12353,23915,29283,38889,14461,36928,3551,38255,12820,43112,40462,19137,37502,43961,18195,44582,18442,39757,11421,40721,30003,1344,1948,35499,6104,34855,8841,37024,3381,9841,8560,48179,5085,17647,13029,41119,11089,11684,40182,2662,17124,9845,20975,6830,29024,32921,8444,26011,39672,32032,23784,9245,36898,6227,35228,15454,14466,1119,28563,28050,11203,38148,1687,25196,18561,18650,33838,20939,44397,48242,20966,38457,21653,4440,28153,46198,46230,14148,5634,5325,25875,35839,38268,19427,38901,43255,39971,45008,25783,8607,9643,3970,38340,20390,38379,19383,1797,9223,27112,37520,30510,49503,28380,18459,33133,16879,8574,38979,41935,46711,2536,35121,5152,41335,21114,45761,19189,28266,27994,3452,302,8087,34629,29592,25123,35185,6587,40470,14907,44234,33749,11166,12986,16475,40924,6173,1500,29182,37714,46857,12893,48978,48613,18512,27139,12675,34691,44340,30269,19095,3833,47571,17763,8350,36531,16055,17860,1764,42402,30922,32707,26742,34445,33747,33336,35048,40757,49432,10561,18586,34435,49210,6268,18034,3760,16274,31431,45052,35289,17098,10931,4432,5669,46750,22095,37967,19078,15370,41017,43435,48783,11815,9517,48912,27582,30904,524,8553,35783,25589,4652,46764,31112,38,7174,10309,40480,7643,48700,6108,30427,15537,4078,41241,49269,30382,42848,31488,43182,24479,30398,29069,29039,5858,29471,9371,21961,25709,5592,5571,32329,23027,40965,9950,22188,45036,25480,32856,16876,1776,45339,43129,29786,27551,11091,17438,23775,36548,21992,37545,8615,19336,21393,15873,2990,46719,13087,25841,35783,31967,37962,26210,24694,18543,2819,29516,34043,4407,23189,2415,40227,15700,41583,10450,41047,12192,34067,1639,9227,8802,43510,37229,33836,38796,13366,7942,32988,26784,37377,4020,10113,5242,20091,86,33207,21066,21745,34923,28994,40520,40131,18112,49396,22508,13060,39880,1371,2271,9772,16512,12700,12757,32448,40332,32662,9089,25626,28683,26233,31090,18267,16974,2061,19428,36414,40055,25572,19077,37021,45831,31655,21885,13706,37382,15845,49130,29237,48861,15512,18769,40482,15281,42981,41175,22168,7763,35350,31896,32553,36245,44274,37866,18317,41680,15116,11146,17456,30905,2029,31930,61876,4871,23169,480956,40958,216176,368383,350531,827523,321496,640529,226427,918099,592680,863274,656639,322990,319477,163975,535923,431247,773820,139059,155737,272982,406050,8297,2617,6360,460,6857,2576,2211,1513,6437,3994,91,517,2712,9090,3319,6342,9910,1535,9301,509,4670,1219}, false,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFirstDegree(tt.args.db, tt.args.investorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirstDegree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstDegree() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMutualConn(t *testing.T) {
	a := App{}
	a.InitializeAndRun()
	defer a.DBClose()
	type args struct {
		db              *sql.DB
		investorId      string
		otherInvestorId string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// localhost:8083/investors/3511/mutual/428
		{"test1", args{a.DB, "3511", "428"}, []int{996,1014,1015,1980,2153,2432,2856,5319,5335,5879,6923,7068,7966,8844,9353,9917,10047,10403,11810,12733,13735,16791,17814,19504,26291,29019,30239,31462,32925,42565,43511,45319,47370,82450}, false,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMutualConn(tt.args.db, tt.args.investorId, tt.args.otherInvestorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMutualConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMutualConn() got = %v, want %v", got, tt.want)
			}
		})
	}
}