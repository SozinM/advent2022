package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FS struct {
	size   int
	parent *FS
	files  map[string]*FS
}

func main() {
	root := &FS{size: 0, parent: nil, files: make(map[string]*FS)}
	maindir := root
	input := "$ ls\n149291 cgc.vzv\ndir cmcrzdt\ndir hwdvrrp\n26925 hwqvsl\ndir lsmv\ndir ngfllcq\ndir ngnzzmpc\ndir pwhjps\ndir rgwnzttf\n260556 tcglclw.hsn\ndir trvznjhb\ndir wgcqrc\n68873 whpnhm\n$ cd cmcrzdt\n$ ls\ndir chqllfw\n95243 hjpf\n108868 hwqvsl\n115004 jpppczvz.mtp\ndir lnsgfnbr\ndir pdtjlb\ndir rqfzvwts\ndir trvznjhb\n$ cd chqllfw\n$ ls\n56623 cgs.hbt\n134804 zqb.grc\n$ cd ..\n$ cd lnsgfnbr\n$ ls\ndir jtzw\ndir ngfllcq\ndir sdm\ndir wlsg\n$ cd jtzw\n$ ls\ndir nfz\n$ cd nfz\n$ ls\n255427 hwqvsl\n94147 tmnjbqq.fzh\n$ cd ..\n$ cd ..\n$ cd ngfllcq\n$ ls\n110661 cdgqtwcv.lzn\n208050 dpf\n$ cd ..\n$ cd sdm\n$ ls\ndir dhwm\ndir ngfllcq\n125983 rfdz.vqm\n24227 tzn\n41909 tzn.vnr\ndir zdzq\n$ cd dhwm\n$ ls\ndir clr\ndir lhv\ndir ncvmgn\n212499 ngfllcq.dcr\n191108 nggnj\ndir pdtjlb\ndir pwhjps\ndir sqqbthdp\ndir trvznjhb\n$ cd clr\n$ ls\ndir lnbc\n87079 npwhncc\n109530 pfqhpr.tzl\n249566 tmnjbqq.fzh\ndir zgmgztcd\n$ cd lnbc\n$ ls\n62635 ftshngp.vbj\n$ cd ..\n$ cd zgmgztcd\n$ ls\n149111 pwhjps.fjm\n$ cd ..\n$ cd ..\n$ cd lhv\n$ ls\ndir phzfwl\n$ cd phzfwl\n$ ls\ndir gmzcjzm\n$ cd gmzcjzm\n$ ls\n263161 vsptqdv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ncvmgn\n$ ls\n245840 hjpf\n21272 pbcjtbg\ndir stm\n$ cd stm\n$ ls\ndir hnbrd\n$ cd hnbrd\n$ ls\n102906 lftjtq.gdt\n45082 vsptqdv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd pdtjlb\n$ ls\n171382 hwqvsl\n$ cd ..\n$ cd pwhjps\n$ ls\n75342 cgc.vzv\n185458 hwqvsl\n254359 ngfllcq.jzd\ndir pdtjlb\n200999 slnlws.sgh\n91174 vqqbcb\ndir zmc\n$ cd pdtjlb\n$ ls\n39001 ngfllcq\n$ cd ..\n$ cd zmc\n$ ls\ndir cjqmc\n$ cd cjqmc\n$ ls\n257668 ctsfdprp\n889 mzr.hnp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd sqqbthdp\n$ ls\n154217 pwhjps.jtn\n$ cd ..\n$ cd trvznjhb\n$ ls\n105431 hwqvsl\n$ cd ..\n$ cd ..\n$ cd ngfllcq\n$ ls\n183850 dcppzj.lmm\n131039 hbpn.zlp\n110398 hjpf\ndir pwhjps\n251784 rqgslj.sqg\ndir szqf\n150728 vsptqdv\n$ cd pwhjps\n$ ls\ndir pzrtwv\n156616 rpbh.ftj\ndir tzn\n$ cd pzrtwv\n$ ls\n197890 tzn\n$ cd ..\n$ cd tzn\n$ ls\n121296 pdtjlb.nmg\n$ cd ..\n$ cd ..\n$ cd szqf\n$ ls\ndir ngfllcq\ndir qtrhn\n$ cd ngfllcq\n$ ls\ndir vnfcczg\n$ cd vnfcczg\n$ ls\n86691 cgc.vzv\n189290 hjds.lwf\n230265 hwqvsl\ndir jbmvmzn\n223129 ngfllcq.mdw\ndir rpcbpjvm\n215119 tmnjbqq.fzh\n$ cd jbmvmzn\n$ ls\ndir flrszsrr\n175047 hjpf\ndir jrzf\ndir ngfllcq\n$ cd flrszsrr\n$ ls\n163007 zdvfmqr.pfq\n$ cd ..\n$ cd jrzf\n$ ls\n32641 qbnz\n$ cd ..\n$ cd ngfllcq\n$ ls\ndir dlcvcd\ndir gcpftfm\n183962 tzn.mjh\n$ cd dlcvcd\n$ ls\n260612 mhf\n$ cd ..\n$ cd gcpftfm\n$ ls\n227489 hwqvsl\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd rpcbpjvm\n$ ls\ndir tnwzgrvw\n$ cd tnwzgrvw\n$ ls\ndir trvznjhb\n$ cd trvznjhb\n$ ls\n127767 pdtjlb.qjw\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd qtrhn\n$ ls\n81716 dngdll\n76367 tdj\n180116 tmnjbqq.fzh\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd zdzq\n$ ls\n192339 bqcmzm.vzw\ndir cplvj\ndir drpmlmf\n92165 mcthl.dzw\ndir qccnln\n$ cd cplvj\n$ ls\ndir gmqddf\n71720 hjpf\n220700 hwqvsl\n260353 lgw.msr\n$ cd gmqddf\n$ ls\n36587 dmdgjrs\n$ cd ..\n$ cd ..\n$ cd drpmlmf\n$ ls\n4896 hjpf\ndir ngfllcq\n65712 pwhjps.mng\n67097 tmnjbqq.fzh\n$ cd ngfllcq\n$ ls\n248967 swvf.prt\n$ cd ..\n$ cd ..\n$ cd qccnln\n$ ls\ndir tzn\n$ cd tzn\n$ ls\n81833 vsptqdv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd wlsg\n$ ls\n181214 nmglzcds.hcg\n195698 pdtjlb.vbr\ndir trvznjhb\n77162 vsptqdv\n$ cd trvznjhb\n$ ls\n237856 trvznjhb\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd pdtjlb\n$ ls\n101237 hwqvsl\ndir jssl\ndir ngfllcq\ndir pvlqvdrw\ndir pwhjps\ndir tzn\n115255 vsptqdv\n$ cd jssl\n$ ls\n76335 smmjjrp\n$ cd ..\n$ cd ngfllcq\n$ ls\n102639 cgc.vzv\n55243 fjfhdtr.ltc\n$ cd ..\n$ cd pvlqvdrw\n$ ls\n38570 pwhjps.qgq\n191322 scnbjgg.gww\n$ cd ..\n$ cd pwhjps\n$ ls\ndir ghfwwj\ndir vtr\n$ cd ghfwwj\n$ ls\n109461 mdtp.ztw\n$ cd ..\n$ cd vtr\n$ ls\ndir fmtpdc\n$ cd fmtpdc\n$ ls\n42101 bcdbqcs.lhp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd tzn\n$ ls\ndir pdtjlb\n$ cd pdtjlb\n$ ls\n154813 vsptqdv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd rqfzvwts\n$ ls\n250112 hwqvsl\n63241 tzn\n149460 tzn.pph\n$ cd ..\n$ cd trvznjhb\n$ ls\n784 bdptcbl.ntt\n108339 cgc.vzv\ndir hnpdrdsm\ndir mnnwcmd\ndir ngfllcq\ndir pbsnd\ndir pdtjlb\n261083 rhl.cjh\ndir trvznjhb\n103421 wjftfs\n$ cd hnpdrdsm\n$ ls\n253895 pwhjps.nps\n36928 ssjhl\n235679 tmnjbqq.fzh\n38049 trvznjhb.dzs\n$ cd ..\n$ cd mnnwcmd\n$ ls\n218411 fvzhntq.rwm\n78694 gwlcbbtm\n243270 hjpf\n15610 trvznjhb.wdj\n$ cd ..\n$ cd ngfllcq\n$ ls\ndir fmhlq\ndir fwbdttbj\ndir hccstwh\n$ cd fmhlq\n$ ls\n142240 rbrwv.hjl\ndir tjpwvb\n256604 tmnjbqq.fzh\ndir trvznjhb\n$ cd tjpwvb\n$ ls\n83436 sfrpt\n$ cd ..\n$ cd trvznjhb\n$ ls\n44433 trvznjhb\n$ cd ..\n$ cd ..\n$ cd fwbdttbj\n$ ls\n4127 hwqvsl\n$ cd ..\n$ cd hccstwh\n$ ls\ndir cbd\n243765 lvzsrqlw.llc\n$ cd cbd\n$ ls\n158372 vzgtjqbd.tmd\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd pbsnd\n$ ls\n172548 cgc.vzv\n261836 pwhjps\n$ cd ..\n$ cd pdtjlb\n$ ls\n73184 cgc.vzv\ndir mdbjwvh\ndir mgt\ndir nbrvvghc\n98702 ngf.gtb\ndir ngfllcq\n224788 pdtjlb\n191754 tmnjbqq.fzh\n189444 tnqwbmzm.vlq\ndir tzn\ndir ztzsg\n$ cd mdbjwvh\n$ ls\n135602 nvt.rjh\n$ cd ..\n$ cd mgt\n$ ls\ndir pbrtf\ndir whflvwv\n$ cd pbrtf\n$ ls\ndir tnnnllg\n$ cd tnnnllg\n$ ls\n269436 hjpf\n$ cd ..\n$ cd ..\n$ cd whflvwv\n$ ls\n185562 ngfllcq\ndir rdl\n150984 trvznjhb\n$ cd rdl\n$ ls\n231952 mqjcttf\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd nbrvvghc\n$ ls\ndir pdtjlb\n$ cd pdtjlb\n$ ls\n40347 hjpf\ndir vwl\n$ cd vwl\n$ ls\n124408 tzn.hjw\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ngfllcq\n$ ls\n54894 bvgf\n$ cd ..\n$ cd tzn\n$ ls\n41506 vnhlvqqw.cvd\n$ cd ..\n$ cd ztzsg\n$ ls\n216433 fzsnpr.vrd\ndir grrngq\ndir hcmvbhp\ndir lbmnq\n19985 ngfllcq\ndir pqqjgbvj\ndir zsnggz\n$ cd grrngq\n$ ls\ndir cqcvb\ndir ngfllcq\ndir shrhb\n$ cd cqcvb\n$ ls\n125712 cgc.vzv\n69800 fpszwd\n160009 rbbwsszz\ndir trvznjhb\n33640 tzghd.fjp\n$ cd trvznjhb\n$ ls\ndir gtjfll\ndir mrncfvnp\ndir pwt\ndir trvznjhb\n173974 tzn\n$ cd gtjfll\n$ ls\n126840 cgc.vzv\n180163 fswwvhwn\ndir gnpcbvmt\n122192 ngfllcq.snb\n$ cd gnpcbvmt\n$ ls\n127188 crhvwb.pct\n218972 mwg\n$ cd ..\n$ cd ..\n$ cd mrncfvnp\n$ ls\n216569 phlbdl\n$ cd ..\n$ cd pwt\n$ ls\n119692 tmnjbqq.fzh\n$ cd ..\n$ cd trvznjhb\n$ ls\n203464 nltqsvd.fhc\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ngfllcq\n$ ls\n31226 rdwczp.hfq\n$ cd ..\n$ cd shrhb\n$ ls\n246035 bnbg\n98513 hjpf\n$ cd ..\n$ cd ..\n$ cd hcmvbhp\n$ ls\ndir cbb\n201230 cgc.vzv\ndir grmmpz\n220707 hjpf\ndir jnr\ndir psv\ndir trvznjhb\n134184 trvznjhb.ghp\n228507 vsptqdv\n$ cd cbb\n$ ls\n17897 hwqvsl\n$ cd ..\n$ cd grmmpz\n$ ls\n214171 spsnch.drs\n$ cd ..\n$ cd jnr\n$ ls\n82130 bhjqplbc.rth\n$ cd ..\n$ cd psv\n$ ls\n215898 pwhjps\n$ cd ..\n$ cd trvznjhb\n$ ls\ndir dfcn\n98111 gwt.fmw\n20948 hjpf\ndir pwhjps\ndir rnlrgd\n$ cd dfcn\n$ ls\ndir fqm\n179114 mbcrjb\ndir trvznjhb\ndir vfrrzdb\n$ cd fqm\n$ ls\ndir frgwsrdh\n$ cd frgwsrdh\n$ ls\n142027 hhpwsl\n$ cd ..\n$ cd ..\n$ cd trvznjhb\n$ ls\ndir nbbb\n34253 ngfllcq.src\ndir qgqmmvg\ndir tfgwmlc\n11919 trvznjhb.qgz\ndir tzn\n24383 zvfzhb.dcw\n$ cd nbbb\n$ ls\n260947 pwhjps.bdq\n$ cd ..\n$ cd qgqmmvg\n$ ls\n67028 wjvhq.tdz\n$ cd ..\n$ cd tfgwmlc\n$ ls\ndir rmcgqpb\ndir wdtmdtz\n$ cd rmcgqpb\n$ ls\n263786 nsmcndc.bjs\n$ cd ..\n$ cd wdtmdtz\n$ ls\n37751 lnspfqv.tpp\n$ cd ..\n$ cd ..\n$ cd tzn\n$ ls\n265035 nqsgm.dhm\n$ cd ..\n$ cd ..\n$ cd vfrrzdb\n$ ls\ndir pzbtsnd\ndir srpdb\n$ cd pzbtsnd\n$ ls\n72770 pdtjlb\n$ cd ..\n$ cd srpdb\n$ ls\n231540 dzgsf\ndir ngfllcq\ndir sjv\n$ cd ngfllcq\n$ ls\n26488 cgc.vzv\n195815 dfjss\n119177 lbjtjqr\n$ cd ..\n$ cd sjv\n$ ls\n225677 msgjj\n228113 ngfllcq\n92448 tzn.rbp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd pwhjps\n$ ls\n171613 nzqd.rzh\n$ cd ..\n$ cd rnlrgd\n$ ls\n132964 hjpf\n146636 hwqvsl\n187596 mlrllbbb.wmh\n92535 trvznjhb\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd lbmnq\n$ ls\n142963 qfpjgvll.ncb\n$ cd ..\n$ cd pqqjgbvj\n$ ls\ndir dfhhzwp\n253570 jjbr.cgf\ndir lchljrwb\ndir pdtjlb\n$ cd dfhhzwp\n$ ls\ndir bqp\n$ cd bqp\n$ ls\n122939 hjpf\n$ cd ..\n$ cd ..\n$ cd lchljrwb\n$ ls\n259475 fqmppdtd.tjm\n$ cd ..\n$ cd pdtjlb\n$ ls\n199658 vsptqdv\n$ cd ..\n$ cd ..\n$ cd zsnggz\n$ ls\n117242 hjpf\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd trvznjhb\n$ ls\ndir bgvqct\n160709 dtq\ndir fldlwj\ndir ndq\n221408 tmnjbqq.fzh\n69148 zvfzt.rjm\n$ cd bgvqct\n$ ls\n66024 cmjrmfn.lpt\n40153 fdlvqgn.dbt\n135848 tmnjbqq.fzh\n$ cd ..\n$ cd fldlwj\n$ ls\n172275 ngfllcq.gbb\n$ cd ..\n$ cd ndq\n$ ls\n117311 bbhlcn.qll\ndir dtzmzgw\n123263 hsb\ndir jnthg\n111208 pdtjlb\n200860 pjq\n$ cd dtzmzgw\n$ ls\ndir tzn\n$ cd tzn\n$ ls\n249561 pszf.zcn\n$ cd ..\n$ cd ..\n$ cd jnthg\n$ ls\n17013 pwhjps.gpp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd hwdvrrp\n$ ls\ndir fgvqft\ndir fwc\ndir sfhsb\ndir tzn\ndir wbtf\n$ cd fgvqft\n$ ls\n215710 wzh\n$ cd ..\n$ cd fwc\n$ ls\n184576 dmvqc.tsr\ndir hgznwf\ndir lfjtqz\ndir ngfllcq\n53748 ngfllcq.vgl\ndir pgpvcp\n$ cd hgznwf\n$ ls\n51012 tmnjbqq.fzh\n$ cd ..\n$ cd lfjtqz\n$ ls\n96666 fwttv.qrp\n203264 nhc.lgd\ndir pwhjps\n213570 tzn\n$ cd pwhjps\n$ ls\n69941 frqq.jnv\n136814 pqmsc.dgz\n185821 rww.trv\n$ cd ..\n$ cd ..\n$ cd ngfllcq\n$ ls\n97361 zcw.zrq\n$ cd ..\n$ cd pgpvcp\n$ ls\n834 nwv.mtw\n$ cd ..\n$ cd ..\n$ cd sfhsb\n$ ls\n78513 pdtjlb\n$ cd ..\n$ cd tzn\n$ ls\ndir lpf\n$ cd lpf\n$ ls\n242317 bngfvvgq.ptp\n82304 ngfllcq.qdz\ndir wsvqtcb\n$ cd wsvqtcb\n$ ls\n32176 vrwlnphn.nnv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd wbtf\n$ ls\n53446 jvvdpn\n41661 ngfllcq.vhl\ndir pwhjps\n231151 tzn\n241080 vdzdhdtb.dgj\ndir vlqmz\n$ cd pwhjps\n$ ls\n200296 hdds.lsw\n$ cd ..\n$ cd vlqmz\n$ ls\n166538 pwhjps.mnq\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd lsmv\n$ ls\ndir dtjjv\n87897 hjpf\n216417 hwqvsl\ndir ngfllcq\ndir pdtjlb\ndir qlnlbcdv\n230724 vsptqdv\n177119 vvzvnn\n$ cd dtjjv\n$ ls\n218742 hjpf\n$ cd ..\n$ cd ngfllcq\n$ ls\n38560 cgc.vzv\n257037 fbttg.jlc\n29948 pwhjps.bvj\n1253 trvznjhb.nzl\n241388 tzn.vdb\ndir wlmtj\n$ cd wlmtj\n$ ls\n51957 hjpf\n27480 pwhjps.hgj\ndir qdjfgz\ndir shb\n182077 tclmtwh.wzr\ndir trvznjhb\n103119 twlf.rnl\n31950 tzn.zfm\n$ cd qdjfgz\n$ ls\n238882 hpms.gll\ndir qpbsmmp\n184633 trvznjhb.nsb\n130374 vsptqdv\n$ cd qpbsmmp\n$ ls\n60269 spsbz\n$ cd ..\n$ cd ..\n$ cd shb\n$ ls\n140111 vsptqdv\n$ cd ..\n$ cd trvznjhb\n$ ls\ndir qjqzppj\n$ cd qjqzppj\n$ ls\n203246 hjpf\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd pdtjlb\n$ ls\n41982 hjpf\n245930 hwqvsl\ndir mmnhtmr\n42314 ngfllcq.tcn\n68269 pdtjlb\n103066 vhtjp.grt\n$ cd mmnhtmr\n$ ls\ndir zqjjgvj\n$ cd zqjjgvj\n$ ls\n263209 nvhflpng.ngd\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd qlnlbcdv\n$ ls\ndir jmd\n58921 pdtjlb.mwb\ndir pzjgmm\ndir qqvrvcw\n79958 rrqmn.zwv\n18158 swjpt.trv\ndir trvznjhb\ndir tzn\n92135 zjb.nns\n268795 zspzsb.szp\n$ cd jmd\n$ ls\n137766 pwhjps\n$ cd ..\n$ cd pzjgmm\n$ ls\n1704 tzn.rhf\n66307 tzn.zll\n116623 vrfvctv.clb\n$ cd ..\n$ cd qqvrvcw\n$ ls\n179302 zrqf.fcn\n$ cd ..\n$ cd trvznjhb\n$ ls\n265026 qfzlgccf.hvz\ndir rbbmmcc\n$ cd rbbmmcc\n$ ls\ndir rtr\n$ cd rtr\n$ ls\ndir dtw\n$ cd dtw\n$ ls\n249472 svs.tgj\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd tzn\n$ ls\n80112 pdtjlb.thm\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ngfllcq\n$ ls\n228868 ggcfwgr.mwh\n10205 gztwg.pwz\n136188 hjpf\n141381 hwqvsl\n250522 pdtjlb.dwg\ndir pwhjps\ndir qcwvfl\ndir tzn\ndir zgwcwqr\n$ cd pwhjps\n$ ls\n19881 tmrljtw\n$ cd ..\n$ cd qcwvfl\n$ ls\n63317 fcjsw.jcj\ndir gvvfsq\n272464 lvqc\n148216 nwppjnwc.sdg\n121107 tzn.ppw\ndir vwfb\n$ cd gvvfsq\n$ ls\n80607 jplds.mjz\n$ cd ..\n$ cd vwfb\n$ ls\ndir gtlfdvjz\n$ cd gtlfdvjz\n$ ls\n228623 jbbplpz\ndir shf\n$ cd shf\n$ ls\n120966 cgc.vzv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd tzn\n$ ls\n215528 cgc.vzv\n112331 gtzcl.rzp\n128653 mqd.dcz\ndir ngfllcq\ndir vmfgbzmw\n$ cd ngfllcq\n$ ls\n207193 qchb.hmv\n$ cd ..\n$ cd vmfgbzmw\n$ ls\n16152 vtlgffn\n$ cd ..\n$ cd ..\n$ cd zgwcwqr\n$ ls\ndir pdtjlb\n110033 vsptqdv\n$ cd pdtjlb\n$ ls\n9746 cgc.vzv\n8010 jdvjpps\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ngnzzmpc\n$ ls\n116647 gmsnm\n157873 hwqvsl\n$ cd ..\n$ cd pwhjps\n$ ls\n96321 cgc.vzv\ndir lcds\ndir tzn\n$ cd lcds\n$ ls\n134975 wcfv.gpd\n$ cd ..\n$ cd tzn\n$ ls\n95149 hjpf\n55950 pwhjps.rpq\n166540 tdt.pgw\n236704 trvznjhb.ccn\n$ cd ..\n$ cd ..\n$ cd rgwnzttf\n$ ls\n122721 hjpf\n$ cd ..\n$ cd trvznjhb\n$ ls\n106424 zvqz\n$ cd ..\n$ cd wgcqrc\n$ ls\n87367 hjpf\n63133 lld\n234148 pwhjps.lcr\ndir rjnnz\n19538 tzn\n233765 zlvznnwj\n$ cd rjnnz\n$ ls\n258856 gpgdm"
	cmd := strings.Split(input, "\n")
	for _, line := range cmd {
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					root = root.parent
				} else {
					root = root.files[tokens[2]]
				}
			}
		} else {
			if tokens[0] == "dir" {
				root.files[tokens[1]] = &FS{size: 0, parent: root, files: make(map[string]*FS)}
			} else {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					panic("Not expected structure")
				}
				root.files[tokens[1]] = &FS{size: size, parent: root}
			}
		}
	}
	//fmt.Println(maindir)
	data := traceDirs(maindir)
	sort.Ints(data)
	fmt.Println(data)

}

func traceDirs(directory *FS) []int {
	sizes := make([]int, 0, 5)
	for _, file := range directory.files {
		if file.size == 0 {
			size := calcSize(file)
			if size >= 208860 {
				sizes = append(sizes, size)
			}
			for _, size := range traceDirs(file) {
				sizes = append(sizes, size)
			}

		}
	}
	return sizes
}

func calcSize(directory *FS) int {
	var total = 0
	for _, file := range directory.files {
		if file.size == 0 {
			total += calcSize(file)
		} else {
			total += file.size
		}
	}
	return total
}
