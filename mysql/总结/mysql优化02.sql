--1��SQL����Ż�
	--���mysql�Է�����Ӱ���ر���ȿ�����ѯ��־�����sql��䣬ͨ����sql���ִ�У����ִ��ʱ�䣬���Ƿ���Ҫ�������������Ż�
	show satus;--�鿴������״̬
	show status like 'Com_%';
	show status like 'Com_select%';--�鿴�û���¼����������ִ��select����
	show status like 'Com_insert%';----------------insert....
	show status like 'Com_update%';--.......
	show status like 'Com_delete%';--...............

	show session status;
	show global status;--��ʾ����������������ִ��sql���Ĵ���
	show global status like 'Com_select%';--��ʾ����������������ִ��select���Ĵ���
	show global status like 'Com_insert%';----------------insert....
	show global status like 'Com_update%';--.......
	show global status like 'Com_delete%';--...............
	
	--�鿴innodbִ�е�sql��Ӱ�������(�û���¼����)
	show status like 'innodb_rows%';
	show status like 'connection';--����mysq�Ĵ���(�����ɹ���ʧ�ܴ���)
	show status like 'uptime';--���������������������ڵĹ���ʱ��
	show status like 'slow_queries';--�鿴����ѯ����(����ѯ��Ĭ��ʱ�䣺10��,����10���¼������ѯ��־��ȥ)
	show variables like '%slow%';--�鿴����ѯ�Ƿ���,һ�㽨�鿪��
	
	--sql������(������ѯ)
	 desc select * from tname; --��ͬ��
	 explain select * from tname;
	 --ͨ������Ľ��������Կ���Ӱ����������Ƿ��õ�����
	 
	
	
	
	
	
	
	
	
	
	
	
	
	
	
