function() {
	pendingCallbacks.push({ id: id, args: arguments });
	go._resolveCallbackPromise();
};
