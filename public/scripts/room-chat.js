        var MAX_CHAT_LIST = 7; // 可展示的LI个数
        var HEIGHT_PER_LIST = 40; // 每个LI的高度（像素）
        var DEBUG = false;
    var EDEBUG = true; // 异常情况的调试信息，优先级低于DEBUG
    var ROLL_SPEED = 'fast'; // 聊天窗口滚动速度
    var LEFT_LI_MAX = 10; //允许保留的LI数量
    // 聊天记录的起始时间
        var msgStartTimestamp = new Date().getTime() - (1000 * 60 * 10 * 6);
    var GET_SEC_PER_ONE_TIME = 1; // 没几秒钟拉取一次数据
    var GET_MSG_LOCK = false; // 获取消息（异步）锁
    var NICKNAME = $("#nickname").val();
    var Room = function() {
      
            // 聊天窗口
            var chatViewRun = function() {
              var timer = setInterval( function() {
                  _getChatData();
                  //clearInterval(timer);
              } , 1000 * GET_SEC_PER_ONE_TIME);
            };

            // 获取数据
            var _getChatData = function() {
        //var result = '{"ErrorCode":0,"ErrorMsg":"successful","Data":"\u003cli class=\'chat-list-li\'\u003e\u003cp class=\'text-center chat-time\'\u003e2015-12-23 22:22:08\u003c/p\u003e\u003cp class=\'text-left chat-name\'\u003e\u003cstrong\u003eGenialX\u003c/strong\u003e \u003ci\u003eSaid\u003c/i\u003e: \u003cspan\u003eBaitch....\u003c/span\u003e\u003c/p\u003e\u003chr class=\'chat-line\'/\u003e\u003c/li\u003e"}';
        //var obj = $.parseJSON(result); 
        
      //return content;
        if( GET_MSG_LOCK == false) {
        GET_MSG_LOCK = true; 
          $.get("/api/getdata?roomid=1720&msgstarttimestamp=" + msgStartTimestamp, function(result){
                      if(DEBUG) alert("getdata ok");
          var obj = $.parseJSON(result);
      if(obj.ErrorCode == 0) {
      // success
      if(DEBUG) alert("服务接口api/getdata返回成功");
      var html = obj.Data;
                        _appendChatData(html);
            _rollChatWin();

          if(obj.LastTime != 0) {
            msgStartTimestamp = 1 +obj.LastTime;
            
            } else {
            }
          
        }else {
      // fail
         if(DEBUG) alert("服务接口api/getdata失败");
         
      }
        GET_MSG_LOCK = false;
      
          });
          
        }
            }

            // 追加数据
            var _appendChatData = function(data) {// append 
                    if(DEBUG) alert("当前有" + $(".chat-list-ul li").length + "个li");
                    $(".chat-list-ul").append(data);
                    if(DEBUG) alert("追加一个节点后有" + $(".chat-list-ul li").length + "个li");
            }

            // 删除节点
            var _removeChatData = function() {
                    if(DEBUG) alert("当前有" + $(".chat-list-ul li").length + "个li");
                    if($(".chat-list-ul li").length > LEFT_LI_MAX) {
                      if(DEBUG) alert("当前节点LI的数量大于了允许保留的最大值" + LEFT_LI_MAX + " 删除节点操作");
                      var removeSum = $(".chat-list-ul li").length - LEFT_LI_MAX;
                      var curMT = parseInt($(".chat-list-ul").css("margin-top"));
                      $(".chat-list-ul li:lt(" + ( $(".chat-list-ul li").length - LEFT_LI_MAX ) + ")").remove();
                      // 移位
                      if(DEBUG) alert("移位" + (curMT + removeSum * HEIGHT_PER_LIST) + "px");
                      $(".chat-list-ul").css({"margin-top": (curMT + removeSum * HEIGHT_PER_LIST) + "px"});
                      if(DEBUG) alert("删除多余节点后有" + $(".chat-list-ul li").length + "个li");
                    }

            }

            // 滚动聊天窗口
            var _rollChatWin = function() {
                var chatViewLi = $(".chat-list-ul li");
                var liSum = chatViewLi.length;
                var moveSum = MAX_CHAT_LIST - liSum;
                if(moveSum < 0) {
                    if(DEBUG) alert("move");
                    var curMT = parseInt($(".chat-list-ul").css("margin-top"));
                    if(DEBUG) alert("margin-top" + curMT);
                    var setMT = 0;
                    var setMT = HEIGHT_PER_LIST * moveSum;
                    if(setMT == curMT) {
                      if(DEBUG) alert("no new data append");
                    } else {
                      if(DEBUG) alert("set margin-top to " + setMT  + "px");
                      $(".chat-list-ul").animate({"margin-top" : setMT   + "px"},ROLL_SPEED,function() {
                         _removeChatData();
                      });
                    }
                } else {
                    if(DEBUG) alert("just it");
                }
            }

            var sendRun = function() {
             
            // 绑定键盘事件
            $("#send-input").keydown(function(event){  
                if(event.shiftKey&& event.which == 13) {
                    // shift + enter
                    if(DEBUG) alert("shift + evnet");
                    _inputSentEvent();
                }
            }); 


            $("#send-button").click(function() {
                  if(DEBUG) alert("mouse click");
                  _inputSentEvent();
              });
            }

            var _inputSentEvent = function() {
                    var data = _getSentInput();
          if(DEBUG) alert("input value:" + data);
          _sendInputData(data);
            };


            // 发送input内容
            var _sendInputData = function(content) {
        // 异步
        if(DEBUG) alert("_sendInputData")
                $.post("/api/postdata", { roomid: "1720", name: NICKNAME, content: content },
                   function(data){
           if(DEBUG) alert("send input data ok");
           var obj = $.parseJSON(data);
           if(obj.ErrorCode == 0) {
            
                    _clearSentInput();
          } else {
            
             if(DEBUG || EDEBUG)alert("抱歉，发送消息失败，请联系作者");
          }
                });
            }

            // 清楚input内容
            var _clearSentInput = function() {
              $("#send-input").val("");
            };

            // 获取INPUT内容
            var _getSentInput = function() {
              return $("#send-input").val();
            }
        
            return {
                init : function() {
                    chatViewRun();
                    sendRun();
                }
            }; 
        }();
